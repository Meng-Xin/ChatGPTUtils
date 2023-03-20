package chatNet

import (
	"chatGPT/global"
	"chatGPT/utils"
	gogpt "github.com/sashabaranov/go-openai"
)

// ChatReq Chat聊天通用信息
type ChatReq struct {
	Model  ChatModel                     `json:"model"`   //会话模型
	Role   ChatRole                      `json:"role"`    //会话角色
	ConnId uint32                        `json:"conn_id"` //会话id
	Token  string                        `json:"token"`   //会话Token有就用，没有就默认
	Msg    []gogpt.ChatCompletionMessage `json:"msg"`
}

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

// ChatPropertyTag
type ChatPropertyTag = string

const (
	HistoryMsgTag ChatPropertyTag = "ChatHistory"
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
func GetMsg(chatRes gogpt.ChatCompletionResponse) gogpt.ChatCompletionMessage {
	if len(chatRes.Choices) == 0 {
		return gogpt.ChatCompletionMessage{}
	}
	return chatRes.Choices[0].Message
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
