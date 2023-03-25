package request

import gogpt "github.com/sashabaranov/go-openai"

type AddToScenesRequest struct {
	Token   string `json:"token" example:""`     // Token
	Timeout int64  `json:"timeout" example:"20"` // 但此请求超时时间
	Scenes
}

type ChatToScenesRequest struct {
	ConnId  uint32     `json:"conn_id" example:"1"` // 会话id
	ChatGPT ChatGPTReq `json:"chat_gpt"`            // 聊天信息
}
type PaintToScenesRequest struct {
	ConnId   uint32             `json:"conn_id"`
	Painting gogpt.ImageRequest `json:"painting"` //绘画模型
}

type GetAllScenesSet struct {
}

type GetScenesSetRequest struct {
}

type SetToScenesRequest struct {
	Token   string `json:"token" `  // Token
	Timeout int64  `json:"timeout"` // 但此请求超时时间
	Scenes
}

type DeleteToConnRequest struct {
}

type ChatGPTReq struct {
	Model int          `json:"model" example:"1"` // 模型选择
	Msg   []ChatGPTMsg `json:"msg"`               // 聊天消息
}

type ChatGPTMsg struct {
	Role    string `json:"role" example:"user"`          // 角色[可不填]
	Content string `json:"content" example:"Hello"`      // 具体内容
	Name    string `json:"name" example:"DefaultWindow"` // 本次会话名称
}

type Scenes struct {
	ScenesID int            `json:"scenes_id" example:"1"`
	ChatGPT  SetChatScenes  `json:"chat_gpt"`
	Paint    SetPaintScenes `json:"paint"`
}

type SetChatScenes struct {
	Model int    `json:"model" example:"1"`            // 会话模型
	Role  string `json:"role" example:"user"`          // 会话角色
	Name  string `json:"name" example:"DefaultWindow"` // 会话标题
}

type SetPaintScenes struct {
	Size           string `json:"size" example:"255x255"`        // 绘画尺寸
	ResponseFormat string `json:"response_format" example:"url"` // 绘画相应格式
	N              int    `json:"n" example:"2"`                 // 绘画数量
}
