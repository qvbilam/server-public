package global

import (
	userProto "file/api/qvbilam/user/v1"
	"file/config"
	"gorm.io/gorm"
)

var (
	DB               *gorm.DB
	ServerConfig     config.ServerConfig
	UserServerClient userProto.UserClient
)
