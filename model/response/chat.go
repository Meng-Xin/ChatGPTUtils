package response

type AddToScenesResponse struct {
	ConnId uint32 `json:"conn_id"` // 场景ID
}
type GetToConnResponse struct {
}
type SetToConnResponse struct {
}
type DeleteToConnResponse struct {
}

type GetConn struct {
}

type ConnSet struct {
	ConnId  uint32 // 链接ID
	Timeout int64  // 单次请求超时时间
	Token   string // 自定义Token
	Scenes  Scenes // 场景id
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
