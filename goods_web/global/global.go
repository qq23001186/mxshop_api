package global

import (
	ut "github.com/go-playground/universal-translator"

	"web_api/goods_web/config"
	"web_api/goods_web/proto"
)

var (
	Trans          ut.Translator
	ServerConfig   *config.ServerConfig = &config.ServerConfig{}
	GoodsSrvClient proto.GoodsClient

	NacosConfig *config.NacosConfig = &config.NacosConfig{}
)
