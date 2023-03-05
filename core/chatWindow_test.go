package core

import (
	"chatGPT/initialize"
	"testing"
)

func TestAddChatWindow(t *testing.T) {
	// 加载初始化文件
	initialize.GlobalInit()
	chatObj := NewChatInfo()
	chatObj.Dialog = "你知道常用的游戏算法有那些么？"
	chatObj.AddChatWindow()
	chatObj.Dialog = "第三种算法怎么实现的？"
	chatObj.AddChatWindow()
}
