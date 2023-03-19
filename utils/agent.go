package utils

import (
	openai "github.com/sashabaranov/go-gpt3"
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
		IdleConnTimeout: 168 * time.Duration(time.Hour),
	}
	config.HTTPClient = &http.Client{
		Transport: transport,
	}
	return config
}
