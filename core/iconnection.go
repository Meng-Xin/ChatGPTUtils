package core

import (
	gogpt "github.com/sashabaranov/go-openai"
	"net"
)

type IConnection interface {
	// Start 	启动链接，当前链接开始工作
	Start()
	// Stop		停止链接，结束当前链接的工作
	Stop()
	// GetConn	获取当前链接绑定的conn
	GetConn() *gogpt.Client
	// GetConnID	获取当前连接ID
	GetConnID() uint32
	// RemoteAddr 	获取客户端状态
	RemoteAddr() net.Addr
	// SetProperty 	设置连接属性
	SetProperty(string, interface{})
	// GetProperty	获取连接属性
	GetProperty(string) (interface{}, error)
	// RemoveProperty 移除连接属性
	RemoveProperty(string)
	// SendMsg 		发送消息
	SendMsg(interface{}) (interface{}, error)
}
