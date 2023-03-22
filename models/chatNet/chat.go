package chatNet

import (
	"chatGPT/global"
	gogpt "github.com/sashabaranov/go-openai"
	"net/http"
	"net/url"
	"time"
)

// PublicProper 通用属性
type PublicProper struct {
	ChatGPT  ChatGPTModel       `json:"chat_gpt"`  //通用聊天模型
	Painting gogpt.ImageRequest `json:"painting"`  //绘画模型
	ConnId   uint32             `json:"conn_id"`   //会话id
	ScenesId int                `json:"scenes_id"` //场景Id
	Token    string             `json:"token"`     //会话Token有就用，没有就默认
	Timeout  int64              `json:"timeout"`   //但此请求超时时间
}

// ChatGPTModel ChatGPT 聊天模型
type ChatGPTModel struct {
	Model ChatModel                     `json:"model"` //会话模型
	Role  ChatRole                      `json:"role"`  //会话角色
	Name  string                        `json:"name"`  //会话别名
	Msg   []gogpt.ChatCompletionMessage `json:"msg"`
}

// ScenesType 场景类型
type ScenesType = int

const (
	ChatGPTScenes  ScenesType = 1 //通用聊天模型
	PaintingScenes ScenesType = 2 //DALL-E 2 绘画模型
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
func GetMsg(chatRes interface{}) gogpt.ChatCompletionMessage {
	if val, ok := chatRes.(gogpt.ChatCompletionResponse); ok {
		if len(val.Choices) == 0 {
			return gogpt.ChatCompletionMessage{}
		}
		return val.Choices[0].Message
	}
	return gogpt.ChatCompletionMessage{}
}

// GetProxyConfig 获取代理配置
func GetProxyConfig(token string, reqTimeout int64) gogpt.ClientConfig {
	// 是否存在自定义Token，使用用户Token
	if token != "" {
		return InitOpenAiAgent(token, global.ProxyPath, global.Config.ChatConn.IdleConnTimeout, reqTimeout)
	} else {
		// 使用默认Token
		return global.OpenAiProxy
	}
}

// InitOpenAiAgent 初始化ChatGPT代理配置
func InitOpenAiAgent(token string, proxyPath string, idleConnTimeout, reqTimeout int64) gogpt.ClientConfig {
	config := gogpt.DefaultConfig(token)
	proxyUrl, err := url.Parse(proxyPath)
	if err != nil {
		panic(err)
	}
	transport := &http.Transport{
		Proxy:           http.ProxyURL(proxyUrl),
		IdleConnTimeout: time.Duration(idleConnTimeout) * time.Hour, // 后续配置文件管理
	}
	config.HTTPClient = &http.Client{
		Transport: transport,
		Timeout:   time.Second * time.Duration(reqTimeout),
	}
	return config
}
