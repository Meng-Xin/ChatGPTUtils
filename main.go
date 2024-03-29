package main

import (
	_ "chatGPT/docs"
	"chatGPT/global"
	"chatGPT/initialize"
	"chatGPT/router"
)

func main() {
	// 加载初始化文件
	initialize.GlobalInit()
	// 开启路由
	r := router.NewRouter()
	r.Run(global.Config.Server.DSN())
	// 关闭服务器清空所有
	defer global.ChatConnManager.ClearConn()
}
