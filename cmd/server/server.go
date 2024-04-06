package server

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/app/user/routers"
	config2 "gitlab.eaip.top/gorm-gen-gin-learn-project/config"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/middleware"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	config string
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
	SvrCmd.PersistentFlags().StringVarP(&config, "config", "c", "$HOME/config.toml", "config file")
}

func startServer() {

	slog.Info("Start loading configuration file...")

	config2.ParseConfig(config)

	slog.Info("Starting kobra service...")

	StartMainServer()
}

func StartMainServer() {
	engine := gin.New()
	middleware.InitMiddleware(engine)
	routers.Get(engine)
	srv := &http.Server{
		Addr:    config2.SConfig.Host + ":" + config2.SConfig.Port,
		Handler: engine,
	}
	go func() {
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("start server error!", "reason", err)
		}
	}()
	slog.Info("Server started!Address:" + config2.SConfig.GetServerAddress())
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	slog.Info("Server shutting down...")
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("Server shutdown error!", "reason", err)
	}
	select {
	case <-ctx.Done():
		slog.Error("Server shutdown timeout!")
	}
	slog.Info("Server exiting")

}
