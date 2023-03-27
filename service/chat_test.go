package service

import (
	"chatGPT/initialize"
	"chatGPT/model/request"
	"chatGPT/pkg/e"
	"fmt"
	"testing"
)

func TestAddChatWindow(t *testing.T) {
	// 加载初始化文件
	initialize.GlobalInit()
	chatObj := ChatService{}
	req := request.AddToScenesRequest{
		Token:   "",
		Timeout: 20,
		Scenes: request.Scenes{
			ScenesID: 1,
			ChatGPT: request.SetChatScenes{
				Model: 1,
				Role:  "system",
				Name:  "测试",
			},
		},
	}
	res := chatObj.AddToScenes(req)
	if res.Status != e.SUCCESS {
		t.Error("ChatGPT is not Role Msg")
	} else {
		fmt.Printf("AddToScenes RESP %+v", res)
	}
}
