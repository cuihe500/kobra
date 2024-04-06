package response

import (
	"github.com/gin-gonic/gin"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/tools"
	"net/http"
)

var Default = &response{}

func Error(c *gin.Context, code int32, message string, err error) {
	res := Default.Clone()
	if err != nil {
		res.SetMeg(err.Error())
	}
	if message != "" {
		res.SetMeg(message)
	}
	res.SetTraceId(tools.GenerateMsgIDFromContext(c))
	res.SetCode(code)
	res.SetSuccess(false)
	c.Set("result", res)
	c.Set("status", http.StatusOK)
	c.AbortWithStatusJSON(http.StatusOK, res)
}
