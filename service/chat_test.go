package service

import (
	"chatGPT/initialize"
	"fmt"
	openai "github.com/sashabaranov/go-gpt3"
	"testing"
)

func TestAddChatWindow(t *testing.T) {
	// 加载初始化文件
	initialize.GlobalInit()
	chatObj := &ChatInfo{
		Msg: []openai.ChatCompletionMessage{{"user", "你知道那些游戏算法"}},
	}
	res := chatObj.AddChatWindow()
	if dialog, ok := res.Data.(*openai.ChatCompletionResponse); ok {
		for _, val := range dialog.Choices {
			fmt.Printf("%+v\n", val)
		}
	}

}
