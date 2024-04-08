package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RuntimeEnv struct {
	db     *gorm.DB
	engine *gin.Engine
}

var Env = &RuntimeEnv{}

func (env *RuntimeEnv) SetDB(db *gorm.DB) {
	env.db = db
}

func (env *RuntimeEnv) DB() *gorm.DB {
	return env.db
}

func (env *RuntimeEnv) Engine() *gin.Engine {
	return env.engine
}

func (env *RuntimeEnv) SetEngine(engine *gin.Engine) {
	env.engine = engine
}

func GetRuntimeEnv() (*RuntimeEnv, error) {
	if Env == nil {
		return nil, errors.New("not have any runtime environment")
	}
	if Env.db == nil {
		return nil, errors.New("not have any database connection")
	}
	if Env.engine == nil {
		return nil, errors.New("not have any engine")
	}
	return Env, nil
}
