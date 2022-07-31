package global

import (
	ut "github.com/go-playground/universal-translator"

	"web_api/user_web/config"
	"web_api/user_web/proto"
)

var (
	Trans         ut.Translator
	ServerConfig  *config.ServerConfig = &config.ServerConfig{}
	UserSrvClient proto.UserClient

	NacosConfig *config.NacosConfig = &config.NacosConfig{}
)
