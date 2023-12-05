package main

import (
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

func InitDB(l LoggerV1) *gorm.DB {
	type Config struct {
		DSN string `yaml:"dsn"`
	}
	var cfg = Config{
		DSN: "root:root@tcp(localhost:3306)/webook?charset=utf8&parseTime=True&loc=Local",
	}
	// 看起来，remote 不支持 key 的切割
	err := viper.UnmarshalKey("db", &cfg)

	db, err := gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{
		Logger: glogger.New(gormLoggerFunc(l.Debug), glogger.Config{
			//	慢查询阈值，只有执行时间超过这个阈值，才会使用
			//	SQL 查询必然要求命中索引，最好就是走一次磁盘 IO
			//	一次磁盘IO是不到10ms
			SlowThreshold:             time.Millisecond * 10,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			LogLevel:                  glogger.Info,
		}),
	})
	if err != nil {
		panic(err)
	}

	db = db.Debug()

	//err = dao.InitTable(db)
	//if err != nil {
	//	panic(err)
	//}

	return db
}

type LoggerV1 interface {
	Debug(msg string, args ...Field)
	Info(msg string, args ...Field)
	Warn(msg string, args ...Field)
	Error(msg string, args ...Field)
}

type Field struct {
	Key   string
	Value any
}
type gormLoggerFunc func(msg string, fields ...Field)

func (g gormLoggerFunc) Printf(msg string, args ...interface{}) {
	g(msg, Field{Key: "args", Value: args})
}
