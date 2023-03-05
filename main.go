package main

import (
	"chatGPT/core"
	"chatGPT/initialize"
)

func main() {
	// 加载初始化文件
	initialize.GlobalInit()
	// 开启路由
	chatObj := core.NewChatInfo()
	chatObj.Dialog = "你知道常用的游戏算法有那些么？"
	chatObj.AddChatWindow()
}
