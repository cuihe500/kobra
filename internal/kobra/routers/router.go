package routers

import (
	"github.com/gin-gonic/gin"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/internal/common"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/internal/middleware"
)

var Routers = make([]func(v1 *gin.RouterGroup, auth gin.HandlerFunc), 0)
var NoAuthenticationRouters = make([]func(v1 *gin.RouterGroup), 0)

/**
该包负责规定Group,并且循环将需要的engine和中间件导入，同时注册路径
*/

func SetupRouter(engine *gin.Engine) {
	for _, r := range Routers {
		r(engine.Group(common.APIVersion), middleware.Authentication())
	}
}
func SetupNoAuthenticationRouter(engine *gin.Engine) {
	for _, r := range NoAuthenticationRouters {
		r(engine.Group(common.APIVersion))
	}
}
