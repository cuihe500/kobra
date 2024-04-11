package handles

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/internal/kobra"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/internal/kobra/user/models"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/internal/pkg/response"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/internal/tools"
	"gorm.io/gorm"
	"log/slog"
	"strconv"
	"time"
)

func GetAllUserInfo(c *gin.Context) {
	db := kobra.Env.DB()
	users := &[]models.User{}
	db.Find(users)
	response.Success(c, users, "查询成功！")
	return
}

func GetUserInfoById(c *gin.Context) {
	db := kobra.Env.DB()
	id := c.Param("id")
	if id, err := strconv.Atoi(id); err != nil || id <= 0 {
		response.Error(c, 400, "", nil)
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
	return
}

func Login(c *gin.Context) {
	db := kobra.Env.DB()
	u := &models.User{}
	if err := c.ShouldBind(u); err != nil {
		slog.Error("绑定用户登录信息模型错误！", "reason", err)
		response.Error(c, 400, "", err)
		return
	}
	user := &models.User{}
	if u.Username != "" {
		if err := db.Where("username = ?", u.Username).First(user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Error("通过用户名查找用户账号信息错误！", "reason", err)
			response.Error(c, 500, "", err)
			return
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, 404, "用户不存在！请检查用户名是否正确！", err)
			return
		} else {
			if u.Password == user.Password {
				jwt, err := tools.GenerateNewJwtToken(user.UUID, user.Username)
				if err != nil {
					response.Error(c, 500, "", err)
					return
				}
				redis := kobra.Env.Redis()
				redis.Set(c, user.UUID.String(), jwt, time.Hour*3)
				response.Success(c, jwt, "登录成功！")
				return
			} else {
				response.Error(c, 401, "用户名或密码错误！", nil)
				return
			}
		}
	}
	if u.Email != "" {
		if err := db.Where("email = ?", u.Email).First(user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Error("通过邮箱查找用户账号信息错误！", "reason", err)
			response.Error(c, 500, "", err)
			return
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, 404, "用户不存在！请检查用户名是否正确！", err)
			return
		} else {
			if u.Password == user.Password {
				jwt, err := tools.GenerateNewJwtToken(user.UUID, user.Username)
				if err != nil {
					response.Error(c, 500, "", err)
					return
				}
				response.Success(c, jwt, "登录成功！")
				return
			} else {
				response.Error(c, 401, "用户名或密码错误！", nil)
				return
			}
		}
	}

}

func Register(c *gin.Context) {
	db := kobra.Env.DB()
	user := &models.User{}
	if err := c.ShouldBind(user); err != nil {
		slog.Error("绑定用户注册信息模型错误！", "reason", err)
		response.Error(c, 400, "传入参数错误！", err)
		return
	}
	if err := db.Create(user).Error; err != nil {
		slog.Error("存储新的用户账号信息错误！", "reason", err)
		response.Error(c, 500, "", err)
		return
	}
	response.Success(c, "", "注册成功！")
	return
}
