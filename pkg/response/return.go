package response

import (
	"github.com/gin-gonic/gin"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/tools"
	"net/http"
)

func Error(c *gin.Context, code int32, message string, err error) {
	res := &response{}
	if err != nil {
		res.SetMeg(err.Error())
	}
	if message != "" {
		res.SetMeg(message)
	}
	res.SetTraceId(tools.GenerateMsgIDFromContext(c))
	res.SetCode(code)
	res.SetSuccess(false)
	res.SetData(nil)
	c.Set("result", res)
	c.Set("status", http.StatusOK)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func Success(c *gin.Context, data interface{}, message string) {
	res := &response{}
	res.SetTraceId(tools.GenerateMsgIDFromContext(c))
	res.SetSuccess(true)
	res.SetData(data)
	res.SetCode(200)
	res.SetMeg(message)
	c.Set("result", res)
	c.Set("status", http.StatusOK)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func Other(c *gin.Context, data interface{}, code int32, message string, success bool) {
	res := &response{}
	res.SetTraceId(tools.GenerateMsgIDFromContext(c))
	res.SetSuccess(success)
	res.SetData(data)
	res.SetCode(code)
	res.SetMeg(message)
	c.Set("result", res)
	c.Set("status", http.StatusOK)
	c.AbortWithStatusJSON(http.StatusOK, res)
}
