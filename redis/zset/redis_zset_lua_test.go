package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func GetClient() redis.Cmdable {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	return redisClient
}

type hotList struct {
	Id     string
	Degree string
	Name   string
}

//go:embed zrange_eval.lua
var luaEvalCode string

type RedisLuaTestSuite struct {
	suite.Suite
	rdb redis.Cmdable
}

func TestRedisLua(t *testing.T) {
	suite.Run(t, &RedisLuaTestSuite{})
}

func (s *RedisLuaTestSuite) SetupSuite() {
	s.rdb = GetClient()
}

func (s *RedisLuaTestSuite) TestGetZSet() {
	testCases := []struct {
		name   string
		before func(t *testing.T)
		after  func(t *testing.T)

		bizs []string

		wantErr error
		wantRes map[string][]hotList
	}{
		{
			name: "success case",
			before: func(t *testing.T) {
				s.rdb.ZAdd(context.Background(), "hotlist:biz:videotest:like", redis.Z{
					Score:  120,
					Member: "id1",
				}, redis.Z{
					Score:  34,
					Member: "id2",
				}, redis.Z{
					Score:  67,
					Member: "id3",
				}, redis.Z{
					Score:  220,
					Member: "id4",
				}, redis.Z{
					Score:  180,
					Member: "id5",
				}, redis.Z{
					Score:  108,
					Member: "id6",
				})
				s.rdb.ZAdd(context.Background(), "hotlist:biz:articletest:like", redis.Z{
					Score:  34,
					Member: "id1",
				}, redis.Z{
					Score:  54,
					Member: "id2",
				}, redis.Z{
					Score:  987,
					Member: "id3",
				}, redis.Z{
					Score:  1212,
					Member: "id4",
				}, redis.Z{
					Score:  654,
					Member: "id5",
				}, redis.Z{
					Score:  3,
					Member: "id6",
				})
			},
			after: func(t *testing.T) {
				s.rdb.Del(context.Background(), "hotlist:biz:videotest:like", "hotlist:biz:articletest:like")
			},
			bizs:    []string{"hotlist:biz:videotest:like", "hotlist:biz:articletest:like", "hotlist:biz:hotlisttest:like"},
			wantErr: nil,
			wantRes: map[string][]hotList{
				"hotlist:biz:videotest:like": {
					{
						Id:     "id4",
						Degree: "220",
					},
					{
						Id:     "id5",
						Degree: "180",
					},
					{
						Id:     "id1",
						Degree: "120",
					},
					{
						Id:     "id6",
						Degree: "108",
					},
					{
						Id:     "id3",
						Degree: "67",
					},
					{
						Id:     "id2",
						Degree: "34",
					},
				},
				"hotlist:biz:articletest:like": {
					{
						Id:     "id4",
						Degree: "1212",
					},
					{
						Id:     "id3",
						Degree: "987",
					},
					{
						Id:     "id5",
						Degree: "654",
					},
					{
						Id:     "id2",
						Degree: "54",
					},
					{
						Id:     "id1",
						Degree: "34",
					},
					{
						Id:     "id6",
						Degree: "3",
					},
				},
				"hotlist:biz:hotlisttest:like": {},
			},
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			tc.before(t)
			res, err := ParseZsetToHotList(s.rdb, tc.bizs)
			assert.Equal(t, tc.wantErr, err)

			if !reflect.DeepEqual(tc.wantRes, res) {
				t.Error("res not equal")
			}

			tc.after(t)
		})
	}

}

func ParseZsetToHotList(rdb redis.Cmdable, bizs []string) (map[string][]hotList, error) {
	redisRes, err := rdb.Eval(context.Background(), luaEvalCode, bizs).Result()
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(redisRes)
	if err != nil {
		return nil, err
	}

	var hotListRes [][]string
	err = json.Unmarshal(b, &hotListRes)
	if err != nil {
		return nil, err
	}

	if len(hotListRes) != len(bizs) {
		// add log
		return nil, err
	}

	res := make(map[string][]hotList)
	for i := range hotListRes {
		if len(hotListRes[i])%2 != 0 {
			// add log
			continue
		}

		res[bizs[i]] = make([]hotList, len(hotListRes[i])/2)

		for j := 0; j < len(hotListRes[i])/2; j++ {
			res[bizs[i]][j] = hotList{
				Id:     hotListRes[i][j*2+0],
				Degree: hotListRes[i][j*2+1],
				Name:   "",
			}

		}
	}

	return res, nil
}
