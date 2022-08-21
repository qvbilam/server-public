package global

import (
	userProto "file/api/user/pb"
	"file/config"
	"gorm.io/gorm"
)

var (
	DB               *gorm.DB
	ServerConfig     config.ServerConfig
	UserServerClient userProto.UserClient
)
