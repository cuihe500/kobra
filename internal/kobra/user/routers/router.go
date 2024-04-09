package routers

import "github.com/gin-gonic/gin"

var Routers = make([]func(v1 *gin.RouterGroup), 0)

func SetupRouter(engine *gin.Engine) {
	for _, r := range Routers {
		r(engine.Group("/api/v1"))
	}

}
