package e

var ErrMsg = map[int]string{
	SUCCESS:         "Success",
	ERROR:           "Fail",
	DataFormatError: "数据格式错误",

	ChatGPT_API_Inaccessible:    "ChatGPT 暂时无法访问！",
	ChatGPT_API_Create_Failed:   "聊天出了点问题！",
	ChatGPT_Manager_GetConnFail: "获取会话连接失败",
}

func GetMsg(code int) string {
	return ErrMsg[code]
}
