package global

import (
	"gorm.io/gorm"
	userProto "public/api/qvbilam/user/v1"
	"public/config"
)

var (
	DB               *gorm.DB
	ServerConfig     config.ServerConfig
	UserServerClient userProto.UserClient
)
