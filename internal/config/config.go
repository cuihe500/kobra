package config

import (
	"github.com/spf13/viper"
	"log/slog"
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

type LogLevelConfig struct {
	DefaultLogLevel  string
	DatabaseLogLevel string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type CasbinConfig struct {
	Host              string
	Port              string
	Username          string
	Password          string
	DatabaseName      string
	RBACModelFilePath string
}

var ServerConf ServerConfig
var DatabaseConf DatabaseConfig
var LogLevelConf LogLevelConfig
var RedisConf RedisConfig
var CasbinConf CasbinConfig

func init() {
	viper.AddConfigPath("$HOME/.kobra_config")
	viper.SetConfigName("kobra_config")
	viper.SetConfigType("toml")
	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("server.port", "8000")
	viper.SetDefault("loglevel.default", "info")
	viper.SetDefault("loglevel.database", "info")
	viper.SetDefault("global.mode", "development")
	viper.SetDefault("redis.host", "127.0.0.1")
	viper.SetDefault("redis.port", "6379")
	viper.SetDefault("redis.password", "")
	viper.SetDefault("redis.db", "0")
	viper.SetDefault("database.host", "127.0.0.1")
	viper.SetDefault("database.port", "3306")
	viper.SetDefault("database.username", "root")
	viper.SetDefault("database.password", "123456")
	viper.SetDefault("database.database_name", "users")
	viper.SetDefault("rbac.host", "127.0.0.1")
	viper.SetDefault("rbac.port", "3306")
	viper.SetDefault("rbac.username", "root")
	viper.SetDefault("rbac.password", "123456")
	viper.SetDefault("rbac.database_name", "casbin")
	viper.SetDefault("rbac.rbac_model_filepath", "config/rbac_model.conf")

	err := viper.BindEnv("server_host", "KOBRA_SERVER_HOST")
	if err != nil {
		slog.Error("Bind ENV Variable KOBRA_SERVER_HOST Failed!", "reason", err)
	}
	err1 := viper.BindEnv("server_port", "KOBRA_SERVER_PORT")
	if err != nil {
		slog.Error("Bind ENV Variable KOBRA_SERVER_PORT Failed!", "reason", err1)
	}
	err2 := viper.BindEnv("database.host", "KOBRA_DATABASE_HOST")
	if err2 != nil {
		slog.Error("Bind ENV Variable KOBRA_DATABASE_HOST Failed!", "reason", err2)
	}
	err3 := viper.BindEnv("database.port", "KOBRA_DATABASE_PORT")
	if err3 != nil {
		slog.Error("Bind ENV Variable KOBRA_DATABASE_PORT Failed!", "reason", err3)
	}
	err4 := viper.BindEnv("database.username", "KOBRA_DATABASE_USERNAME")
	if err4 != nil {
		slog.Error("Bind ENV Variable KOBRA_DATABASE_USERNAME Failed!", "reason", err4)
	}
	err5 := viper.BindEnv("database.password", "KOBRA_DATABASE_PASSWORD")
	if err5 != nil {
		slog.Error("Bind ENV Variable KOBRA_DATABASE_PASSWORD Failed!", "reason", err5)
	}
	err6 := viper.BindEnv("database.database_name", "KOBRA_DATABASE_DATABASE_NAME")
	if err6 != nil {
		slog.Error("Bind ENV Variable KOBRA_DATABASE_DATABASE_NAME Failed!", "reason", err6)
	}
	err7 := viper.BindEnv("loglevel.default", "KOBRA_LOGLEVEL_DEFAULT")
	if err7 != nil {
		slog.Error("Bind ENV Variable KOBRA_LOGLEVEL_DEFAULT Failed!", "reason", err7)
	}
	err8 := viper.BindEnv("loglevel.database", "KOBRA_LOGLEVEL_DATABASE")
	if err8 != nil {
		slog.Error("Bind ENV Variable KOBRA_LOGLEVEL_DATABASE Failed!", "reason", err8)
	}
	err9 := viper.BindEnv("redis.host", "KOBRA_REDIS_HOST")
	if err9 != nil {
		slog.Error("Bind ENV Variable KOBRA_REDIS_HOST Failed!", "reason", err9)
	}
	err10 := viper.BindEnv("redis.port", "KOBRA_REDIS_PORT")
	if err10 != nil {
		slog.Error("Bind ENV Variable KOBRA_REDIS_PORT Failed!", "reason", err10)
	}
	err11 := viper.BindEnv("redis.password", "KOBRA_REDIS_PASSWORD")
	if err11 != nil {
		slog.Error("Bind ENV Variable KOBRA_REDIS_PASSWORD Failed!", "reason", err11)
	}
	err12 := viper.BindEnv("redis.db", "KOBRA_REDIS_DB")
	if err12 != nil {
		slog.Error("Bind ENV Variable KOBRA_REDIS_DB Failed!", "reason", err12)
	}
	err13 := viper.BindEnv("rbac.host", "KOBRA_RBAC_HOST")
	if err13 != nil {
		slog.Error("Bind ENV Variable KOBRA_RBAC_HOST Failed!", "reason", err13)
	}
	err14 := viper.BindEnv("rbac.port", "KOBRA_RBAC_PORT")
	if err14 != nil {
		slog.Error("Bind ENV Variable KOBRA_RBAC_PORT Failed!", "reason", err14)
	}
	err15 := viper.BindEnv("rbac.username", "KOBRA_RBAC_USERNAME")
	if err15 != nil {
		slog.Error("Bind ENV Variable KOBRA_RBAC_USERNAME Failed!", "reason", err15)
	}
	err16 := viper.BindEnv("rbac.password", "KOBRA_RBAC_PASSWORD")
	if err16 != nil {
		slog.Error("Bind ENV Variable KOBRA_RBAC_PASSWORD Failed!", "reason", err16)
	}
	err17 := viper.BindEnv("rbac.database_name", "KOBRA_RBAC_DATABASE_NAME")
	if err17 != nil {
		slog.Error("Bind ENV Variable KOBRA_RBAC_DATABASE_NAME Failed!", "reason", err17)
	}
	err18 := viper.BindEnv("rbac.rbac_model_filepath", "KOBRA_RBAC_MODEL_FILEPATH")
	if err18 != nil {
		slog.Error("Bind ENV Variable KOBRA_RBAC_MODEL_FILEPATH Failed!", "reason", err18)
	}
}

func (config ServerConfig) GetServerAddress() string {
	return "http://" + config.Host + ":" + config.Port
}

func (config DatabaseConfig) GetDSN() string {
	return config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.DatabaseName + "?charset=utf8mb4&parseTime=True&loc=Local"
}
func (config CasbinConfig) GetRbacDSN() string {
	return config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.DatabaseName + "?charset=utf8mb4&parseTime=True&loc=Local"
}
