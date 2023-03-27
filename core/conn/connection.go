package conn

import (
	"bytes"
	"chatGPT/core/idefine"
	"chatGPT/global"
	"chatGPT/model"
	"chatGPT/model/request"
	"chatGPT/utils"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	gogpt "github.com/sashabaranov/go-openai"
	log "github.com/sirupsen/logrus"
	"image/png"
	"io"
	"net"
	"net/http"
	"os"
	"sync"
)

/*
	1.控制chat链接管理
*/

type ChatConnection struct {
	// 当前Conn属于那个用户
	BelongPerson model.User
	// 当前ChatGPT Socket TCP 套接字
	Conn *gogpt.Client
	// 获取 ConnID
	ConnID uint32
	// 当前连接上下文[创建链接时指定]
	Ctx context.Context
	// 当前会话场景
	Scenes *Scenes

	// HttpClient gogpt库小写的Conn导致无法获取到http链接指针，只能自己拿
	client *http.Client
	//  当前链接状态
	isClosed bool
	// 链接属性集合
	property map[string]interface{}
	// 保护连接的锁
	propertyLock sync.RWMutex
}

// NewChatConn 创建一个chatGPT connection 实例，connID：链接id，model：GPT模型 role：GPT角色 token：用户自定义Token
func NewChatConn(connId uint32, req request.AddToScenesRequest) *ChatConnection {
	connConfig := GetProxyConfig(req.Token, req.Timeout)
	c := &ChatConnection{
		Conn:     gogpt.NewClientWithConfig(connConfig),
		ConnID:   connId,
		Ctx:      context.Background(),
		Scenes:   &Scenes{},
		client:   connConfig.HTTPClient,
		property: make(map[string]interface{}),
	}
	err := c.Scenes.SetScenes(req.Scenes)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	// 初始聊天记录
	c.property[HistoryMsgTag] = make([]gogpt.ChatCompletionMessage, 0)
	// 添加到管理模块
	global.ChatConnManager.Add(c)
	return c
}

func (c *ChatConnection) GetScenes() idefine.IScenes {
	return c.Scenes
}

func (c *ChatConnection) Start() {
	//TODO 使用链接调用函数
	panic("implement me")
}

func (c *ChatConnection) Stop() {
	// 关闭Http链接
	c.isClosed = true
	c.client.CloseIdleConnections()
}

func (c *ChatConnection) GetConn() *gogpt.Client {
	return c.Conn
}

func (c *ChatConnection) GetConnID() uint32 {
	return c.ConnID
}

func (c *ChatConnection) RemoteAddr() net.Addr {
	//TODO 获取客户端地址
	panic("implement me")
}

func (c *ChatConnection) SetProperty(key string, val interface{}) {
	// 设置对话链接属性
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()
	c.property[key] = val
}

func (c *ChatConnection) GetProperty(key string) (interface{}, error) {
	// 获取对话链接属性
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()
	if val, ok := c.property[key]; ok {
		return val, nil
	}
	return nil, errors.New("no property found")
}

func (c *ChatConnection) RemoveProperty(key string) {
	// 删除对话链接属性
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()
	delete(c.property, key)
}

func (c *ChatConnection) SendMsg(req interface{}) (resp interface{}, err error) {
	if c.isClosed {
		return resp, errors.New("The httpconnection is closed")
	}
	reqData, ok := req.(request.ChatGPTReq)
	if !ok {
		return nil, errors.New("SendMsg 入参不是 PublicProper")
	}
	switch c.Scenes.ScenesID {
	case ChatGPTScenes:
		return c.sendToChatGPT(reqData.Msg)
	case PaintingScenes:
		//return c.sendToDALL(reqData.Painting)
	}

	return resp, errors.New("ChatConn is not found Scenes！")
}

func (c *ChatConnection) sendToChatGPT(val []request.ChatGPTMsg) (interface{}, error) {
	var historyMsg []gogpt.ChatCompletionMessage
	data := make([]gogpt.ChatCompletionMessage, len(val))
	// 结构体属性相同直接强转
	for i, v := range val {
		data[i] = gogpt.ChatCompletionMessage(v)
	}
	// 替换角色
	for k, datum := range data {
		if datum.Role == "" {
			data[k].Role = c.Scenes.ChatGPT.Role
		}
	}
	// 获取历史消息并保存玩家对话
	propVal, err := c.GetProperty(HistoryMsgTag)
	if err != nil {
		return gogpt.ChatCompletionResponse{}, err
	}
	historyMsg, ok := propVal.([]gogpt.ChatCompletionMessage)
	if !ok {
		return gogpt.ChatCompletionResponse{}, errors.New("conn property not found HistoryMsgTag!")
	}
	historyMsg = append(historyMsg, data...)
	resp, err := c.Conn.CreateChatCompletion(
		c.Ctx,
		gogpt.ChatCompletionRequest{
			Model:    gogpt.GPT3Dot5Turbo,
			Messages: historyMsg,
		},
	)
	if err != nil {
		return resp, err
	}
	// 保存Ai对话
	aiMsg := utils.GetMsg(resp)
	historyMsg = append(historyMsg, aiMsg)
	c.SetProperty(HistoryMsgTag, historyMsg)
	return resp, nil
}

func (c *ChatConnection) sendToDALL(data gogpt.ImageRequest) (interface{}, error) {
	switch data.ResponseFormat {
	case gogpt.CreateImageResponseFormatURL:
		// 构建请求
		respUrl, err := c.Conn.CreateImage(c.Ctx, data)
		if err != nil {
			return nil, err
		}
		log.Info(respUrl.Data[0].URL)
		return respUrl.Data[0].URL, nil
	case gogpt.CreateImageResponseFormatB64JSON:
		respBase64, err := c.Conn.CreateImage(c.Ctx, data)
		if err != nil {
			return nil, err
		}
		imgBytes, err := base64.StdEncoding.DecodeString(respBase64.Data[0].B64JSON)
		if err != nil {
			fmt.Printf("Base64 decode error: %v\n", err)
			return nil, err
		}
		r := bytes.NewReader(imgBytes)
		imgData, err := png.Decode(r)
		if err != nil {
			return nil, err
		}

		file, err := os.Create("image.png")
		if err != nil {
			fmt.Printf("File creation error: %v\n", err)
			return nil, err
		}
		defer file.Close()

		if err := png.Encode(file, imgData); err != nil {
			fmt.Printf("PNG encode error: %v\n", err)
			return nil, err
		}
		return respBase64.Data[0].B64JSON, nil
	}
	return nil, errors.New("未匹配到绘画模型参数！")
}

func (c *ChatConnection) SendMsgToChatStream(data []gogpt.ChatCompletionMessage) error {
	if c.isClosed {
		return errors.New("The httpconnection is closed")
	}
	// 替换角色
	for k, datum := range data {
		if datum.Role == "" {
			data[k].Role = c.Scenes.ChatGPT.Role
		}
	}

	resp, err := c.Conn.CreateChatCompletionStream(
		c.Ctx,
		gogpt.ChatCompletionRequest{
			Model:    gogpt.GPT3Dot5Turbo,
			Messages: data,
			Stream:   true,
		},
	)
	defer resp.Close()
	if err != nil {
		fmt.Println("CreateChatCompletionStream Failed error:", err)
		return nil
	}
	fmt.Printf("Stream response: ")
	for {
		response, err := resp.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			return nil
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return err
		}
		//msgChan <- response.Choices[0].Delta.Content
		fmt.Println(response.Choices[0].Delta.Content)
	}

}
