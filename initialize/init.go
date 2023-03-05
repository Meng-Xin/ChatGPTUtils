package initialize

import (
	"chatGPT/config"
	"chatGPT/global"
	"chatGPT/utils"
)

func GlobalInit() {
	// 初始化全局配置文件
	global.Config = config.InitLoadConfig()
	// 是否开启代理,代理配置初始化
	if global.Config.Server.OpenProxy {
		global.OpenAiProxy = utils.InitOpenAiAgent(global.OpenAiToken, global.ProxyPath)
	}
	// 中间件初始化

	// 雪花id
	global.SnowId = utils.NewWorker(1, 1)
}
