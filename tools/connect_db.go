package tools

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := "root:Ch870219176!@tcp(192.168.125.112:43307)/users?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		StdoutLog.Error("Connect DB fail!", "reason", err)
	}

}
