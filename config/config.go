package config

import (
	"github.com/spf13/viper"
	"log/slog"
	"os"
)

type ServerConfig struct {
	Host string
	Port string
}

type DatabaseConfig struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DatabaseName string
}

var SConfig ServerConfig
var DConfig DatabaseConfig

func init() {
	viper.AddConfigPath("$HOME/.kobra_config")
	viper.SetConfigName("kobra_config")
	viper.SetConfigType("toml")
	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("server.port", "8000")

	err := viper.BindEnv("KOBRA_SERVER_HOST")
	if err != nil {
		slog.Error("Bind ENV Variable KOBRA_SERVER_HOST Failed!", "reason", err)
	}
	err1 := viper.BindEnv("KOBRA_SERVER_PORT")
	if err != nil {
		slog.Error("Bind ENV Variable KOBRA_SERVER_PORT Failed!", "reason", err1)
	}
	err2 := viper.BindEnv("KOBRA_DATABASE_HOST")
	if err2 != nil {
		slog.Error("Bind ENV Variable KOBRA_DATABASE_HOST Failed!", "reason", err2)
	}
	err3 := viper.BindEnv("KOBRA_DATABASE_PORT")
	if err3 != nil {
		slog.Error("Bind ENV Variable KOBRA_DATABASE_PORT Failed!", "reason", err3)
	}
	err4 := viper.BindEnv("KOBRA_DATABASE_USERNAME")
	if err4 != nil {
		slog.Error("Bind ENV Variable KOBRA_DATABASE_USERNAME Failed!", "reason", err4)
	}
	err5 := viper.BindEnv("KOBRA_DATABASE_PASSWORD")
	if err5 != nil {
		slog.Error("Bind ENV Variable KOBRA_DATABASE_PASSWORD Failed!", "reason", err5)
	}
	err6 := viper.BindEnv("KOBRA_DATABASE_DATABASE_NAME")
	if err6 != nil {
		slog.Error("Bind ENV Variable KOBRA_DATABASE_DATABASE_NAME Failed!", "reason", err6)
	}
}

func ParseConfig(path string) {

	viper.SetConfigFile(path)

	err := viper.ReadInConfig()
	if err != nil {
		slog.Error("Read config file failed!", "reason", err)
		os.Exit(1)
	}

	SConfig = ServerConfig{
		Host: viper.GetString("server.host"),
		Port: viper.GetString("server.port"),
	}
	DConfig = DatabaseConfig{
		Host:         viper.GetString("database.host"),
		Port:         viper.GetString("database.port"),
		Username:     viper.GetString("database.username"),
		Password:     viper.GetString("database.password"),
		DatabaseName: viper.GetString("database.database_name"),
	}

}
func (config ServerConfig) GetServerAddress() string {
	return "http://" + config.Host + ":" + config.Port
}

func (config DatabaseConfig) GetDSN() string {
	return config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.DatabaseName + "?charset=utf8mb4&parseTime=True&loc=Local"
}