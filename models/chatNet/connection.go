package chatNet

import (
	"chatGPT/global"
	"chatGPT/models"
	"context"
	"errors"
	gogpt "github.com/sashabaranov/go-gpt3"
	"net"
	"sync"
)

/*
	1.控制chat链接管理
*/

type ChatConnection struct {
	// 当前Conn属于那个用户
	BelongPerson models.User
	// 当前ChatGPT Socket TCP 套接字
	Conn *gogpt.Client
	// 获取 ConnID
	ConnID uint32
	// 当前连接上下文[创建链接时指定]
	Ctx context.Context

	//  当前链接状态
	isClosed bool
	// 当前链接模型[创建链接时指定]
	model string
	// [创建链接时指定] 角色 ai：聊天对象为ai human：聊天对象为正常人类 agent：聊天对象为代理
	role string
	// 链接属性集合
	property map[string]interface{}
	// 保护连接的锁
	propertyLock sync.RWMutex
}

// NewChatConn 创建一个chatGPT connection 实例，connID：链接id，model：GPT模型 role：GPT角色
func NewChatConn(connId uint32, model ChatModel, role ChatRole, token string) *ChatConnection {
	c := &ChatConnection{
		Conn:     gogpt.NewClientWithConfig(GetProxyConfig(token)),
		ConnID:   connId,
		Ctx:      context.Background(),
		model:    SwitchGPTModel(model),
		role:     SwitchGPTRole(role),
		property: make(map[string]interface{}),
	}
	// 添加到管理模块
	global.ChatConnManager.Add(c)
	return c
}

func (c ChatConnection) Start() {
	//TODO 使用链接调用函数
	panic("implement me")
}

func (c ChatConnection) Stop() {
	//TODO 链接关闭，执行回调函数。
	panic("implement me")
}

func (c ChatConnection) GetConn() *gogpt.Client {
	return c.Conn
}

func (c ChatConnection) GetConnID() uint32 {
	return c.ConnID
}

func (c ChatConnection) RemoteAddr() net.Addr {
	//TODO 获取客户端地址
	panic("implement me")
}

func (c ChatConnection) SendMsg(data []gogpt.ChatCompletionMessage) (gogpt.ChatCompletionResponse, error) {
	// 替换角色
	for k, datum := range data {
		if datum.Role == "" {
			data[k].Role = c.role
		}
	}
	resp, err := c.Conn.CreateChatCompletion(
		c.Ctx,
		gogpt.ChatCompletionRequest{
			Model:    gogpt.GPT3Dot5Turbo,
			Messages: data,
		},
	)
	return resp, err
}

func (c ChatConnection) SetProperty(key string, val interface{}) {
	// 设置对话链接属性
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()
	c.property[key] = val
}

func (c ChatConnection) GetProperty(key string) (interface{}, error) {
	// 获取对话链接属性
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()
	if val, ok := c.property[key]; ok {
		return val, nil
	}
	return nil, errors.New("no property found")
}

func (c ChatConnection) RemoveProperty(key string) {
	// 删除对话链接属性
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()
	delete(c.property, key)
}
