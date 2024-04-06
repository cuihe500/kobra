package tools

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log/slog"
)

const TrafficKey = "X-Request-Id"

func GetDBConnect() *gorm.DB {
	dsn := config.DConfig.GetDSN()
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("Connect DB fail!", "reason", err)
	}
	return db
}

func GenerateMsgIDFromContext(ctx *gin.Context) string {
	requestID := ctx.GetHeader(TrafficKey)
	if requestID == "" {
		requestID = uuid.New().String()
		ctx.Header(TrafficKey, requestID)
	}
	return requestID
}
