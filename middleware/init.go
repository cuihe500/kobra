package middleware

import "github.com/gin-gonic/gin"

func InitMiddleware(e *gin.Engine) {
	e.Use(WithContextDb)
}
