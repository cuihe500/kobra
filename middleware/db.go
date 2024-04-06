package middleware

import (
	"github.com/gin-gonic/gin"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/tools"
)

func WithContextDb(ctx *gin.Context) {
	db := tools.GetDBConnect()
	ctx.Set("db", db.WithContext(ctx))
	ctx.Next()
}
