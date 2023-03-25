package global

import (
	"chatGPT/config"
	"chatGPT/core/idefine"
	openai "github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
	"sync"
)

const (
	OpenAiToken = "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" // ChatGPT Token
	ProxyPath   = "http://127.0.0.1:7890"
)

var (
	OpenAiProxy     openai.ClientConfig  // OpenAIProxy 代理配置
	Config          *config.AllConfig    // 全局config
	SourceConnID    *ConnID              // SourceConnID 用于生成ConnID
	ChatConnManager idefine.IConnManager // ChatConnManager 连接管理
	MysqlDB         *gorm.DB             // MysqlDB
)

type ConnID struct {
	SourceConnID uint32
	connIDLock   sync.Mutex
}

func (c *ConnID) GetConnID() uint32 {
	c.connIDLock.Lock()
	defer c.connIDLock.Unlock()
	c.SourceConnID += 1
	return c.SourceConnID
}
