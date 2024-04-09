package routers

import (
	"github.com/gin-gonic/gin"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/internal/kobra/user/handles"
	"net/http"
)

func init() {
	Routers = append(Routers, GetAllUserRouter)
}

func GetAllUserRouter(router *gin.RouterGroup) {
	router.
		GET("/users", handles.GetAllUserInfo).
		GET("/user/get/id/:id", handles.GetUserInfoById).
		GET("/what-i-want-to-say", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Written for love and peace.\nBest wishes to you Xu QianQian.")
		}).
		POST("/login", handles.Login).
		POST("/register", handles.Register)
}