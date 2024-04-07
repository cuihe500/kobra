package handles

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/app"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/app/user/models"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/app/user/models/dto"
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
	return
}

func GetUserInfoById(c *gin.Context) {
	db := app.Env.DB()
	id := c.Param("id")
	if id == "" {
		response.Error(c, 400, "传入参数id为空！", nil)
		return
	}
	if id, err := strconv.Atoi(id); err != nil || id <= 0 {
		response.Error(c, 400, "传入参数错误！", nil)
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
	db := app.Env.DB()
	var userLoginDto = &dto.UserLoginDto{}
	if err := c.ShouldBind(userLoginDto); err != nil ||
		(userLoginDto.Email == "" && userLoginDto.Username == "") ||
		(userLoginDto.Email != "" && userLoginDto.Username != "") ||
		userLoginDto.Password == "" {
		slog.Error("绑定并验证用户登录信息模型错误！", "errorMsg", err)
		response.Error(c, 400, "传入参数错误！", err)
		return
	}
	user := &models.User{}
	if userLoginDto.Username != "" {
		if err := db.Where("username = ?", userLoginDto.Username).First(user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Error("通过用户名查找用户账号信息错误！", "reason", err)
			response.Error(c, 500, "", err)
			return
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, 404, "用户不存在！请检查用户名是否正确！", err)
			return
		} else {
			if user.Password == userLoginDto.Password {
				response.Success(c, "", "登录成功！")
				return
			} else {
				response.Error(c, 401, "用户名或密码错误！", nil)
				return
			}
		}
	}
	if userLoginDto.Email != "" {
		if err := db.Where("email = ?", userLoginDto.Email).First(user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Error("通过邮箱查找用户账号信息错误！", "reason", err)
			response.Error(c, 500, "", err)
			return
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, 404, "用户不存在！请检查用户名是否正确！", err)
			return
		} else {
			if user.Password == userLoginDto.Password {
				response.Success(c, "", "登录成功！")
				return
			} else {
				response.Error(c, 401, "用户名或密码错误！", nil)
				return
			}
		}
	}

}

func Register(c *gin.Context) {
	db := app.Env.DB()
	var userRegisterDto = &dto.UserRegisterDto{}
	if err := c.ShouldBind(userRegisterDto); err != nil ||
		(userRegisterDto.Email == "" || userRegisterDto.Username == "" || userRegisterDto.Password == "") {
		slog.Error("绑定并验证用户注册信息模型错误！", "errorMes", err)
		response.Error(c, 400, "传入参数错误！", err)
		return
	}
	user := (&models.User{}).CreateNewUserModel(userRegisterDto)
	if err := db.Create(user).Error; err != nil {
		slog.Error("存储新的用户账号信息错误！", "reason", err)
		response.Error(c, 500, "", err)
		return
	}
	response.Success(c, "", "注册成功！")
	return
}
