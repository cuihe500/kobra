package tools

import (
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
