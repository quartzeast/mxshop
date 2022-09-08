package global

import (
	"mxshop/services/user/config"

	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	ServerConfig config.Server
	NacosConfig  config.Nacos
)
