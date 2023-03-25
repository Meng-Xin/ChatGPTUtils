package utils

import (
	"encoding/json"
	gogpt "github.com/sashabaranov/go-openai"
)

func JsonMarshal(obj gogpt.ChatCompletionMessage) ([]byte, error) {
	marshal, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

func JsonUnmarshal() {

}

// GetMsg 获取本次对话文本
func GetMsg(chatRes interface{}) gogpt.ChatCompletionMessage {
	val, ok := chatRes.(gogpt.ChatCompletionResponse)
	if !ok || len(val.Choices) == 0 {
		return gogpt.ChatCompletionMessage{}
	}
	return val.Choices[0].Message
}

// GetImages 获取本次绘图信息
func GetImages(paint interface{}) gogpt.ImageResponse {
	val, ok := paint.(gogpt.ImageResponse)
	if !ok || len(val.Data) == 0 {
		return gogpt.ImageResponse{}
	}
	return val
}
