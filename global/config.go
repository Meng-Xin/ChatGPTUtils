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
