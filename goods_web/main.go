package main

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"web_api/goods_web/global"
	"web_api/goods_web/initialize"
	"web_api/goods_web/utils"
)

func main() {
	//1. 初始化logger
	initialize.InitLogger()
	//2. 初始化配置文件
	initialize.InitConfig()
	//3. 初始化routers
	Router := initialize.Routers()
	//4. 初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}

	//5. 初始化srv的连接
	initialize.InitSrvConn()

	viper.AutomaticEnv()
	//如果是本地开发环境端口号固定，线上环境启动获取端口号
	flag := viper.GetInt("DEV_CONFIG")
	if flag > 4 { // 1=zsz  2=comp  3=home
		port, err := utils.GetFreePort()
		if err == nil {
			global.ServerConfig.Port = port
		}
	}

	/*
		1. S()可以获取一个全局的sugar，可以让我们自己设置一个全局的logger
		2. 日志是分级别的，debug， info ， warn， error， fetal
			debug最低，fetal最高，如果配置成info，所有比info低的都不会输出
			NewProduction默认日志级别为info
			NewDevelopment默认日志级别为debug
		3. S函数和L函数很有用， 提供了一个全局的安全访问logger的途径
	*/
	zap.S().Debugf("启动服务器， 端口: %d", global.ServerConfig.Port)

	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败:", err.Error())
	}
}
