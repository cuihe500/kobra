package tools

import (
	"errors"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Enforcer *casbin.Enforcer

func CasbinEnforcer() (*casbin.Enforcer, error) {
	gormLoggerLevel := getDatabaseLogLevel()
	db, err := gorm.Open(mysql.Open(config.CasbinConf.GetRbacDSN()), &gorm.Config{
		Logger: gormLoggerLevel,
	})
	if err != nil {
		return nil, errors.New("初始化Casbin数据库连接失败！")
	}
	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return nil, errors.New("初始化Casbin Adapter失败！")
	}
	enforcer, err := casbin.NewEnforcer("config/rbac_model.conf", a)
	if err != nil {
		return nil, errors.New("初始化Casbin Enforcer失败！")
	}
	return enforcer, nil
}
