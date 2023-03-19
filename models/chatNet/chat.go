package chatNet

import (
	"chatGPT/global"
	"chatGPT/utils"
	gogpt "github.com/sashabaranov/go-gpt3"
)

// ChatModel ChatGPT Model
type ChatModel = int

const (
	GPT3Dot5Turbo0301 ChatModel = 0
	GPT3Dot5Turbo     ChatModel = 1
)

func SwitchGPTModel(model ChatModel) (rely string) {
	switch model {
	case GPT3Dot5Turbo0301:
		rely = gogpt.GPT3Dot5Turbo0301
	case GPT3Dot5Turbo:
		rely = gogpt.GPT3Dot5Turbo
	default:
		rely = gogpt.GPT3Dot5Turbo
	}
	return rely
}

// ChatRole ChatGPT Role
type ChatRole = int

const (
	RoleToSystem    ChatRole = 1 // 系统
	RoleToUser      ChatRole = 2 // 用户
	RoleToAssistant ChatRole = 3 // 助手
	RoleToAi        ChatRole = 4 // Ai
	RoleToHuman     ChatRole = 5 // 人类
	RoleToAgent     ChatRole = 6 // 代理
)

func SwitchGPTRole(role ChatRole) (rely string) {
	switch role {
	case RoleToSystem:
		rely = "system"
	case RoleToUser:
		rely = "user"
	case RoleToAssistant:
		rely = "assistant"
	case RoleToAi:
		rely = "ai"
	case RoleToHuman:
		rely = "human"
	case RoleToAgent:
		rely = "agent"
	default:
		rely = "user"
	}
	return rely
}

// GetMsg 获取本次对话文本
func GetMsg(chatRes gogpt.ChatCompletionResponse) string {
	return chatRes.Choices[0].Message.Content
}

// GetProxyConfig 获取代理配置
func GetProxyConfig(token string) gogpt.ClientConfig {
	// 是否存在自定义Token，使用用户Token
	if token != "" {
		return utils.InitOpenAiAgent(token, global.ProxyPath)
	} else {
		// 使用默认Token
		return global.OpenAiProxy
	}
}
