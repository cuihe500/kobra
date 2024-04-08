package server

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/app"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/app/user/models"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/app/user/routers"
	config2 "gitlab.eaip.top/gorm-gen-gin-learn-project/config"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/tools"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	config string
	merge  bool
	logger *slog.Logger
	SvrCmd = &cobra.Command{
		Use:   "server",
		Long:  "Start a new kobra service.",
		Short: "Start the kobra service.",
		Run: func(cmd *cobra.Command, args []string) {
			startServer()
		},
	}
)

func init() {
	SvrCmd.PersistentFlags().StringVarP(&config, "config", "c", "$HOME/config.toml", "Set config file location.")
	SvrCmd.PersistentFlags().BoolVarP(&merge, "merge", "m", false, "Init database and merge it.")
}

func startServer() {

	slog.Info("Start loading configuration file...")
	parseConfig(config)

	logger = tools.GetDefaultLogger()
	slog.SetDefault(logger)

	slog.Info("Start connect database...")
	initDBConnect()
	slog.Info("Init kobra service...")
	if merge {
		slog.Info("Start init database and merge it...")
		mergeDatabase()
	}
	initGinServer()
	slog.Info("Init Routers...")
	initRouters()
	slog.Info("Start kobra service...")
	startMainServer()
}

func parseConfig(path string) {

	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		slog.Error("Parse config file failed!", "reason", err)
	}

	config2.ServerConf = config2.ServerConfig{
		Host: viper.GetString("server.host"),
		Port: viper.GetString("server.port"),
	}
	config2.DatabaseConf = config2.DatabaseConfig{
		Host:         viper.GetString("database.host"),
		Port:         viper.GetString("database.port"),
		Username:     viper.GetString("database.username"),
		Password:     viper.GetString("database.password"),
		DatabaseName: viper.GetString("database.database_name"),
	}
	config2.LogLevelConf = config2.LogLevelConfig{
		DefaultLogLevel:  viper.GetString("loglevel.default"),
		DatabaseLogLevel: viper.GetString("loglevel.database"),
	}
}

func initGinServer() {
	if viper.GetString("global.mode") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.New()
	engine.Use(sloggin.New(logger))
	app.Env.SetEngine(engine)
}

func initDBConnect() {
	db, err := tools.GetDbConnect()
	if err != nil {
		slog.Error("Create DB Connection Failed!", "reason", err)
		os.Exit(1)
	}
	app.Env.SetDB(db)
}

func initRouters() {
	routers.SetupRouter(app.Env.Engine())
}

func startMainServer() {
	srv := &http.Server{
		Addr:    config2.ServerConf.Host + ":" + config2.ServerConf.Port,
		Handler: app.Env.Engine(),
	}
	go func() {
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("start server error!", "reason", err)
			os.Exit(1)
		}
	}()
	slog.Info("Server started!Address:" + config2.ServerConf.GetServerAddress())
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	slog.Info("Server shutting down...")
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("Server shutdown error!", "reason", err.Error())
	}
	select {
	case <-ctx.Done():
		slog.Error("Server shutdown timeout!")
	}
	slog.Info("Server exiting")

}

func mergeDatabase() {
	db := app.Env.DB()
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		slog.Error("Init and merge database failed!", "reason", err)
		os.Exit(1)
	}
}
