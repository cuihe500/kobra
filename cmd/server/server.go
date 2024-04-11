package server

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	config2 "gitlab.eaip.top/gorm-gen-gin-learn-project/internal/config"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/internal/kobra"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/internal/kobra/user/models"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/internal/kobra/user/routers"
	tools2 "gitlab.eaip.top/gorm-gen-gin-learn-project/internal/tools"
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

	logger = tools2.GetDefaultLogger()
	slog.SetDefault(logger)

	slog.Info("Start connect database...")
	initDBConnect()
	slog.Info("Start connect Redis...")
	initRedisConnect()

	slog.Info("Init RBAC Casbin...")
	initCasbin()

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
		return
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
	config2.RedisConf = config2.RedisConfig{
		Host:     viper.GetString("redis.host"),
		Port:     viper.GetString("redis.port"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	}
	config2.CasbinConf = config2.CasbinConfig{
		Host:              viper.GetString("rbac.host"),
		Port:              viper.GetString("rbac.port"),
		Username:          viper.GetString("rbac.username"),
		Password:          viper.GetString("rbac.password"),
		DatabaseName:      viper.GetString("rbac.database_name"),
		RBACModelFilePath: viper.GetString("rbac.rbac_model_file"),
	}
}

func initGinServer() {
	if viper.GetString("global.mode") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.New()
	engine.Use(sloggin.New(logger))
	kobra.Env.SetEngine(engine)
}

func initDBConnect() {
	db, err := tools2.DbConnect()
	if err != nil {
		slog.Error("Create DB Connection Failed!", "reason", err)
		os.Exit(1)
	}
	kobra.Env.SetDB(db)
}

func initRouters() {
	routers.SetupRouter(kobra.Env.Engine())
	routers.SetupNoAuthenticationRouter(kobra.Env.Engine())
}

func startMainServer() {
	srv := &http.Server{
		Addr:    config2.ServerConf.Host + ":" + config2.ServerConf.Port,
		Handler: kobra.Env.Engine(),
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
	db := kobra.Env.DB()
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		slog.Error("Init and merge database failed!", "reason", err)
		os.Exit(1)
	}
}
func initRedisConnect() {
	rdb, err := tools2.RedisConnect()
	if err != nil {
		slog.Error("Init Redis Connection Failed!", "reason", err)
		os.Exit(1)
	}
	kobra.Env.SetRedis(rdb)
}
func initCasbin() {
	var err error
	tools2.Enforcer, err = tools2.CasbinEnforcer()
	if err != nil {
		slog.Error("初始化Casbin失败！", "reason", err)
	}
}
