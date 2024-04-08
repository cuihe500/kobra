package tools

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	slogGorm "github.com/orandin/slog-gorm"
	"github.com/phsym/console-slog"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"log/slog"
	"os"
)

const TrafficKey = "X-Request-Id"

func GetDbConnect() (*gorm.DB, error) {
	dsn := config.DatabaseConf.GetDSN()
	gormLoggerLevel := getDatabaseLogLevel()
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLoggerLevel,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GenerateMsgIDFromContext(ctx *gin.Context) string {
	requestID := ctx.GetHeader(TrafficKey)
	if requestID == "" {
		requestID = uuid.New().String()
		ctx.Header(TrafficKey, requestID)
	}
	return requestID
}

func getDatabaseLogLevel() gormLogger.Interface {
	if config.LogLevelConf.DatabaseLogLevel == "silent" {
		return slogGorm.New().LogMode(gormLogger.Silent)
	}
	if config.LogLevelConf.DatabaseLogLevel == "info" {
		return slogGorm.New().LogMode(gormLogger.Info)
	}
	if config.LogLevelConf.DatabaseLogLevel == "warn" {
		return slogGorm.New().LogMode(gormLogger.Warn)
	}
	if config.LogLevelConf.DatabaseLogLevel == "error" {
		return slogGorm.New().LogMode(gormLogger.Error)
	}
	return slogGorm.New().LogMode(gormLogger.Info)

}
func GetDefaultLogger() *slog.Logger {
	if config.LogLevelConf.DefaultLogLevel == "debug" {
		return slog.New(
			console.NewHandler(os.Stdout, &console.HandlerOptions{Level: slog.LevelDebug}),
		)
	}
	if config.LogLevelConf.DefaultLogLevel == "info" {
		return slog.New(
			console.NewHandler(os.Stdout, &console.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	if config.LogLevelConf.DefaultLogLevel == "warn" {
		return slog.New(
			console.NewHandler(os.Stdout, &console.HandlerOptions{Level: slog.LevelWarn}),
		)
	}
	if config.LogLevelConf.DefaultLogLevel == "error" {
		return slog.New(
			console.NewHandler(os.Stdout, &console.HandlerOptions{Level: slog.LevelError}),
		)
	}
	return slog.New(
		console.NewHandler(os.Stdout, &console.HandlerOptions{Level: slog.LevelInfo}),
	)
}
