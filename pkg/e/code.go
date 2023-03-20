package e

// HTTP Code
const (
	SUCCESS         = 200 // Success
	ERROR           = 400 // Fail
	DataFormatError = 505 // 数据格式错误
)

// Business Code

const (
	ChatGPT_API_Inaccessible    = 2000 // ChatGPT 暂时无法访问！
	ChatGPT_API_Create_Failed   = 2001 // 创建窗口失败
	ChatGPT_Manager_GetConnFail = 2002 // 获取会话连接失败
	ChatGPT_Manager_StreamFail  = 2003 // StreamFail
)
