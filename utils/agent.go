package utils

import (
	"chatGPT/global"
	openai "github.com/sashabaranov/go-openai"
	"net/http"
	"net/url"
	"time"
)

func InitOpenAiAgent(token string, proxyPath string) openai.ClientConfig {
	config := openai.DefaultConfig(token)
	proxyUrl, err := url.Parse(proxyPath)
	if err != nil {
		panic(err)
	}
	transport := &http.Transport{
		Proxy:           http.ProxyURL(proxyUrl),
		IdleConnTimeout: time.Duration(global.Config.ChatConn.IdleConnTimeout) * time.Hour, // 后续配置文件管理
	}
	config.HTTPClient = &http.Client{
		Transport: transport,
		Timeout:   time.Second * time.Duration(global.Config.ChatConn.Timeout),
	}
	return config
}
