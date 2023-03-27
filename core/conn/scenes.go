package conn

import (
	"chatGPT/model/request"
	"errors"
)

type Scenes struct {
	ScenesID int       // 场景ID
	ChatGPT  SChatGPT  // 聊天模型
	Painting SPainting // 绘画模型
}

// SChatGPT 通用聊天模型
type SChatGPT struct {
	Model string // 当前链接模型[创建链接时指定]
	Role  string // [创建链接时指定] 角色 ai：聊天对象为ai human：聊天对象为正常人类 agent：聊天对象为代理
	Name  string // 会话名称
}

// SPainting DALL-E 2 image generation
type SPainting struct {
	Size           string // 绘画尺寸
	ResponseFormat string // 绘画相应格式
	N              int    // 绘画数量
}

func (s *Scenes) GetScenesID() int {
	return s.ScenesID
}

func (s *Scenes) GetScene() *Scenes {
	return s
}

func (s *Scenes) SetScenes(i interface{}) error {
	data, ok := i.(request.Scenes)
	if !ok {
		return errors.New("assertion request.SetScenes type Fail")
	}
	s.ScenesID = data.ScenesID
	switch s.ScenesID {
	case 1:
		chat := data.ChatGPT
		s.ChatGPT = SChatGPT{Model: SwitchGPTModel(chat.Model), Role: chat.Role, Name: chat.Name}
		return nil
	case 2:
		paint := data.Paint
		s.Painting = SPainting{Size: paint.Size, ResponseFormat: paint.ResponseFormat, N: paint.N}
		return nil
	}
	return errors.New("Not Found Scenes")
}
