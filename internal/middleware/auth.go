package middleware

import (
	"github.com/gin-gonic/gin"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/internal/kobra"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/internal/pkg/response"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/internal/tools"
	"log/slog"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("authorization")
		if header == "" {
			slog.Info("403认证失败！", "reason", "认证请求头为空！", "requestID", c.GetHeader("X-Request-Id"))
			response.Error(c, 403, "用户未登录！", nil)
			return
		}
		token, err := tools.ValidateAndParseJwtToken(header)
		if err != nil {
			slog.Info("Token认证失败！", "reason", err, "requestID", c.GetHeader("X-Request-Id"))
			response.Error(c, 403, "Token验证失败！", nil)
			return
		}
		uuid := token.UUID
		if uuid == "" {
			slog.Info("403认证失败！", "reason", "未能通过Token获取到用户UUID", "requestID", c.GetHeader("X-Request-Id"))
			response.Error(c, 403, "未能通过Token获取到用户UUID", nil)
			return
		}
		rdb := kobra.Env.Redis()
		val := rdb.Get(c, uuid).Val()
		if val == "" {
			slog.Info("403认证失败！", "reason", "未能通过UUID查找到用户信息", "requestID", c.GetHeader("X-Request-Id"))
			response.Error(c, 403, "未能通过UUID查找到用户信息！", nil)
			return
		}
		if val != header {
			slog.Info("403认证失败！", "reason", "Token非法！", "requestID", c.GetHeader("X-Request-Id"))
			response.Error(c, 403, "该Token为非法Token！", nil)
			return
		}
		ok, err := enforce(uuid, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			response.Error(c, 500, "验证权限信息错误！", err)
			return
		}
		if !ok {
			response.Error(c, 403, "该用户无权限访问！", nil)
			return
		}
		c.Next()
	}
}

func enforce(sub string, obj string, act string) (bool, error) {
	slog.Info("验证权限", "sub", sub, "obj", obj, "act", act)
	err := tools.Enforcer.LoadPolicy()
	if err != nil {
		slog.Error("读取权限列表失败！", "reason", err)
		return false, err
	}
	ok, err := tools.Enforcer.Enforce(sub, obj, act)
	if err != nil {
		slog.Error("验证权限失败！", "reason", err)
		return false, err
	}
	return ok, nil
}
