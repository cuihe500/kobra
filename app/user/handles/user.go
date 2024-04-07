package handles

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/app"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/app/user/models"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/pkg/response"
	"gorm.io/gorm"
	"log/slog"
	"strconv"
)

func GetAllUserInfo(c *gin.Context) {
	db := app.Env.DB()
	users := &[]models.User{}
	db.Find(users)
	response.Success(c, users, "查询成功！")
}

func GetUserInfoById(c *gin.Context) {
	db := app.Env.DB()
	id := c.Param("id")
	if id == "" {
		response.Error(c, 401, "传入参数id为空！", nil)
		return
	}
	if id, err := strconv.Atoi(id); err != nil || id <= 0 {
		response.Error(c, 401, "传入参数错误！", nil)
		return
	}
	user := &models.User{}
	if err := db.Where("id = ?", id).First(user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error("通过id查询用户信息错误！", "reason", err)
		response.Error(c, 500, "", err)
		return
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		response.Error(c, 404, "用户未找到！", err)
		return
	}
	response.Success(c, user, "查询成功！")
}
