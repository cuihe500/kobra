package tools

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDbConnect() (*gorm.DB, error) {
	dsn := config.DatabaseConf.GetDSN()
	gormLoggerLevel := getDatabaseLogLevel()
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLoggerLevel,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetRedisConnect() (*redis.Client, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisConf.Host + ":" + config.RedisConf.Port,
		Password: config.RedisConf.Password,
		DB:       config.RedisConf.DB,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return rdb, nil
}
