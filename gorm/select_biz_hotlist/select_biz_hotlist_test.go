package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type GormTestSuite struct {
	suite.Suite
	db     *gorm.DB
	client redis.Cmdable
}

func GetRedisClient() redis.Cmdable {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	return redisClient
}

func TestGorm(t *testing.T) {
	suite.Run(t, &GormTestSuite{})
}

func (s *GormTestSuite) SetupSuite() {
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	s.db = InitDB(NewZapLogger(l))
	s.db.AutoMigrate(&InteractiveTest{})

	s.client = GetRedisClient()
}

func (s *GormTestSuite) TestDistinct() {
	testCases := []struct {
		name   string
		before func(t *testing.T)
		after  func(t *testing.T)

		bizs []string

		wantErr error
	}{
		{
			name: "测试distinct",
			before: func(t *testing.T) {
				//s.db.Raw("source data/video_data.sql")
				//s.db.Raw("source data/article_data.sql")
			},
			bizs: []string{"article", "video"},
			after: func(t *testing.T) {
				//s.db.Exec("delete from interactive_tests")
			},
		},
	}
	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			tc.before(t)

			var bizs []string
			s.db.Model(&InteractiveTest{}).Distinct().Pluck("biz", &bizs)
			assert.Equal(t, bizs, tc.bizs)

			for i := range bizs {
				hotList := getHotListByBiz(s.db, bizs[i])
				fmt.Println(hotList)
				saveHotListToCache(bizs[i], hotList, s.client)
			}

			tc.after(t)
		})
	}
}

type HotList struct {
	Biz  string
	Id   string
	Cnt  string
	Name string
}

func getHotListByBiz(db *gorm.DB, biz string) []InteractiveTest {
	var res []InteractiveTest
	db.Model(&InteractiveTest{}).Where("biz = ? and like_cnt > 0", biz).Order("like_cnt DESC").Limit(100).Scan(&res)
	return res
}

func saveHotListToCache(biz string, list []InteractiveTest, client redis.Cmdable) {
	zset := make([]redis.Z, len(list))
	for i := range list {
		zset[i] = redis.Z{
			Score:  float64(list[i].LikeCnt),
			Member: list[i].BizId,
		}
	}
	client.ZAdd(context.Background(), fmt.Sprintf("hotlist:biz:%s:like", biz), zset...)
}

//	SELECT biz, biz_id, like_cnt FROM `interactives` where biz = 'article' and like_cnt > 0 ORDER BY like_cnt DESC LIMIT 100

type InteractiveTest struct {
	Id         int64  `gorm:"primaryKey,autoIncrement"`
	BizId      int64  `gorm:"uniqueIndex:biz_id_type"`
	Biz        string `gorm:"uniqueIndex:biz_id_type;type:varchar(128)"`
	ReadCnt    int64
	LikeCnt    int64
	CollectCnt int64
	Ctime      int64
	Utime      int64
}
