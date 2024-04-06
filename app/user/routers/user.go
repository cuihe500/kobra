package routers

import (
	"github.com/gin-gonic/gin"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/app/user/handles"
)

var userHandle handles.UserHandle

func Get(engine *gin.Engine) {
	engine.GET("/users/", userHandle.GetAllUserInfo)
}
