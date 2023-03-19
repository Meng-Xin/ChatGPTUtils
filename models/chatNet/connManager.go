package chatNet

import (
	"chatGPT/core"
	"errors"
	"sync"
)

type ConnManager struct {
	// 链接管理集合
	chatConnections map[uint32]core.IConnection
	// 链接管理读写锁
	chatConnLock sync.RWMutex
}

func NewChatConnManager() *ConnManager {
	return &ConnManager{
		chatConnections: make(map[uint32]core.IConnection),
	}
}

// Add 添加一个链接到管理模块中
func (c ConnManager) Add(conn core.IConnection) {
	c.chatConnLock.Lock()
	defer c.chatConnLock.Unlock()
	c.chatConnections[conn.GetConnID()] = conn
}

// Remove 从管理模块中删除对应链接
func (c ConnManager) Remove(conn core.IConnection) {
	c.chatConnLock.Lock()
	defer c.chatConnLock.Unlock()
	delete(c.chatConnections, conn.GetConnID())
}

// Get 从管理模块中获取ConnID对应链接
func (c ConnManager) Get(connID uint32) (core.IConnection, error) {
	c.chatConnLock.RLock()
	defer c.chatConnLock.RUnlock()
	if conn, ok := c.chatConnections[connID]; ok {
		return conn, nil
	}
	return nil, errors.New("chatConnection not Found ")
}

// Len 统计当前连接总量
func (c ConnManager) Len() int {
	c.chatConnLock.RLock()
	defer c.chatConnLock.RUnlock()
	return len(c.chatConnections)
}

// ClearConn 清除连接管理模块下的所有链接
func (c ConnManager) ClearConn() {
	//TODO
	c.chatConnLock.Lock()
	defer c.chatConnLock.Unlock()
	for connID, conn := range c.chatConnections {
		// 停止函数调用
		conn.Stop()
		// 删除对应链接
		delete(c.chatConnections, connID)
	}
}
