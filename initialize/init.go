package initialize

import (
	"chatGPT/config"
	chatNet2 "chatGPT/core/conn"
	"chatGPT/global"
)

func GlobalInit() {
	// 初始化全局配置文件
	global.Config = config.InitLoadConfig()

	// 是否开启代理,代理配置初始化
	if global.Config.Server.OpenProxy {
		global.OpenAiProxy = chatNet2.InitOpenAiAgent(global.OpenAiToken, global.ProxyPath, global.Config.ChatConn.IdleConnTimeout, global.Config.ChatConn.Timeout)
	}
	// 数据库引擎初始化
	InitDatabase(global.Config.Mysql.Dsn(), "")
	// 中间件初始化

	// ChatID Init
	global.SourceConnID = &global.ConnID{}
	// ChatConnectionManager Init
	global.ChatConnManager = chatNet2.NewChatConnManager()
}
