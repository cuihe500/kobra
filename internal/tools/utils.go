package tools

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const TrafficKey = "X-Request-Id"

func GenerateMsgIDFromContext(ctx *gin.Context) string {
	requestID := ctx.GetHeader(TrafficKey)
	if requestID == "" {
		requestID = uuid.New().String()
		ctx.Header(TrafficKey, requestID)
	}
	return requestID
}
