package core

import (
	"chatGPT/global"
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-gpt3"
)

type ChatInfo struct {
	snowID uint64
	Dialog string `json:"dialog"` // 对话信息
	Role   string `json:"role"`   // 对话角色 默认user
}

func NewChatInfo() *ChatInfo {
	snowId, err := global.SnowId.NextID()
	if err != nil {
		fmt.Errorf("Generate snowId failed error：%s", err)
		return nil
	}
	return &ChatInfo{
		snowID: snowId,
		Role:   "user",
	}
}

// AddChatWindow 创建对话窗口
func (c *ChatInfo) AddChatWindow() {
	client := openai.NewClientWithConfig(global.OpenAiProxy)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    c.Role,
					Content: c.Dialog,
				},
			},
		},
	)
	// 如果接口存在错误那么输出错误日志，但是仍然保留对话
	if err != nil {
		fmt.Errorf("ChatGPT Api fail error：%s", err)
	} else {
		for _, choice := range resp.Choices {
			fmt.Println(choice.Message)
		}
	}
}

// RemoveChatWindow

// GetChatWindow

// SetChatWindow
