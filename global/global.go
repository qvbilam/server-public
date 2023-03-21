package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	userProto "public/api/qvbilam/user/v1"
	"public/config"
)

var (
	DB               *gorm.DB
	Redis            redis.Client
	ServerConfig     config.ServerConfig
	UserServerClient userProto.UserClient
)
