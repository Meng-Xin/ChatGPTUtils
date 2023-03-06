package service

import (
	"chatGPT/global"
	"chatGPT/pkg/e"
	"chatGPT/pkg/public"
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-gpt3"
)

type ChatReq struct {
	snowID uint64
	Msg    []openai.ChatCompletionMessage `json:"msg"`
}

// AddChatWindow 创建对话窗口
func (c *ChatReq) AddChatWindow() public.Response {
	code := e.SUCCESS
	client := openai.NewClientWithConfig(global.OpenAiProxy)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: c.Msg,
		},
	)
	// 如果接口存在错误那么输出错误日志，但是仍然保留对话
	if err != nil {
		code = e.ChatGPT_API_Inaccessible
		fmt.Printf("ChatGPT Api fail error：%s\n", err)
		return public.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return public.Response{
		Status: code,
		Data:   resp,
		Msg:    e.GetMsg(code),
	}
}

// RemoveChatWindow

// GetChatWindow

// SetChatWindow

// PingChatMsg 打印输出信息
func PingChatMsg(msg []openai.ChatCompletionChoice) {
	for _, val := range msg {
		fmt.Printf("+v", val)
	}
}
