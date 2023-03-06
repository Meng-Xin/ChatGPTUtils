package global

import (
	"chatGPT/config"
	"chatGPT/utils"
	openai "github.com/sashabaranov/go-gpt3"
)

const (
	OpenAiToken = "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" // ChatGPT Token
	ProxyPath   = "http://127.0.0.1:7890"
)

var (
	OpenAiProxy openai.ClientConfig // OpenAIProxy 代理配置
	Config      *config.AllConfig   // 全局config
	SnowId      *utils.Worker       // 雪花id
)
