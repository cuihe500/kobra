package handles

import (
	"github.com/gin-gonic/gin"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/app"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/app/user/models"
	"net/http"
)

type UserHandle struct {
	app.Handle
}

func (userHandles *UserHandle) GetAllUserInfo(c *gin.Context) {
	orm := userHandles.MakeContext(c).MakeOrm().Orm
	var users []models.User
	orm.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
