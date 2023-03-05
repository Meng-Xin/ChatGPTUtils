package utils

import (
	openai "github.com/sashabaranov/go-gpt3"
	"net/http"
	"net/url"
)

func InitOpenAiAgent(token string, proxyPath string) openai.ClientConfig {
	config := openai.DefaultConfig(token)
	proxyUrl, err := url.Parse(proxyPath)
	if err != nil {
		panic(err)
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	config.HTTPClient = &http.Client{
		Transport: transport,
	}
	return config
}
