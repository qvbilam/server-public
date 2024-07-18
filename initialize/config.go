package initialize

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"public/global"
	"strconv"
)

func InitConfig() {
	initEnvConfig()
	initViperConfig()
}

func initEnvConfig() {
	serverPort, _ := strconv.Atoi(os.Getenv("PORT"))
	DBPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	RedisPort, _ := strconv.Atoi(os.Getenv("REDIS_PORT"))
	RedisDB, _ := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
	userServerPort, _ := strconv.Atoi(os.Getenv("USER_SERVER_PORT"))

	// server
	global.ServerConfig.Name = os.Getenv("SERVER_NAME")
	global.ServerConfig.Port = serverPort
	// database
	global.ServerConfig.DBConfig.Host = os.Getenv("DB_HOST")
	global.ServerConfig.DBConfig.Port = DBPort
	global.ServerConfig.DBConfig.User = os.Getenv("DB_USER")
	global.ServerConfig.DBConfig.Password = os.Getenv("DB_PASSWORD")
	global.ServerConfig.DBConfig.Database = os.Getenv("DB_DATABASE")
	// redis
	global.ServerConfig.RedisConfig.Host = os.Getenv("REDIS_HOST")
	global.ServerConfig.RedisConfig.Port = RedisPort
	global.ServerConfig.RedisConfig.User = os.Getenv("REDIS_USER")
	global.ServerConfig.RedisConfig.Password = os.Getenv("REDIS_PASSWORD")
	global.ServerConfig.RedisConfig.Database = RedisDB
	// user-server
	global.ServerConfig.UserServerConfig.Name = os.Getenv("USER_SERVER_NAME")
	global.ServerConfig.UserServerConfig.Host = os.Getenv("USER_SERVER_HOST")
	global.ServerConfig.UserServerConfig.Port = int64(userServerPort)
}

func initViperConfig() {
	file := "config.yaml"
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return
	}

	v := viper.New()
	v.SetConfigFile(file)
	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		zap.S().Panicf("获取配置异常: %s", err)
	}
	// 映射配置文件
	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		zap.S().Panicf("加载配置异常: %s", err)
	}
	// 动态监听配置
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&global.ServerConfig)
	})
}
