package tools

import (
	slogGorm "github.com/orandin/slog-gorm"
	"github.com/phsym/console-slog"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/internal/config"
	gormLogger "gorm.io/gorm/logger"
	"log/slog"
	"os"
)

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
