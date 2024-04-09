package routers

import (
	"github.com/gin-gonic/gin"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/internal/kobra/user/handles"
	"net/http"
)

func init() {
	Routers = append(Routers, SetRouter)
	NoAuthenticationRouters = append(NoAuthenticationRouters, SetNoAuthenticationRouter)
}

func SetNoAuthenticationRouter(router *gin.RouterGroup) {
	router.
		GET("/what-i-want-to-say", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Written for love and peace.\nBest wishes to you Xu QianQian.")
		}).
		POST("/login", handles.Login).
		POST("/register", handles.Register)
}

func SetRouter(router *gin.RouterGroup, auth gin.HandlerFunc) {
	router.
		GET("/users", auth, handles.GetAllUserInfo).
		GET("/user/get/id/:id", auth, handles.GetUserInfoById)
}
