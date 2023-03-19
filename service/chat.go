package service

import (
	"chatGPT/global"
	"chatGPT/models/chatNet"
	"chatGPT/pkg/e"
	"chatGPT/pkg/public"
	"fmt"
	openai "github.com/sashabaranov/go-gpt3"
)

type ChatService struct {
	ChatReq chatNet.ChatReq `json:"chat_req"`
}

// AddChatWindow 创建对话窗口
func (c *ChatService) AddChatWindow() public.Response {
	code := e.SUCCESS
	// 使用ChatConnManager 进行调用
	conn := chatNet.NewChatConn(global.SourceConnID.GetConnID(), c.ChatReq)
	if conn == nil {
		code = e.ChatGPT_API_Create_Failed
		return public.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return public.Response{
		Status: code,
		Data:   conn.GetConnID(),
		Msg:    e.GetMsg(code),
	}
}

// RemoveChatWindow remove chatGPT Windows And remove chatConn[需要校验删除用户信息]
func (c *ChatService) RemoveChatWindow() public.Response {
	code := e.SUCCESS
	// 删除对应conn
	conn, err := global.ChatConnManager.Get(c.ChatReq.ConnId)
	if err != nil {
		code = e.ChatGPT_Manager_GetConnFail
		return public.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	global.ChatConnManager.Remove(conn)
	return public.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// GetChatWindow get the created link and send msg[需要校验删除用户信息]
func (c *ChatService) GetChatWindow() public.Response {
	code := e.SUCCESS
	// 获取已创建的会话连接
	conn, err := global.ChatConnManager.Get(c.ChatReq.ConnId)
	if err != nil {
		code = e.ChatGPT_Manager_GetConnFail
		return public.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 发送消息
	resData, err := conn.SendMsg(c.ChatReq.Msg)
	resMsg := chatNet.GetMsg(resData)
	if err != nil {
		code = e.ChatGPT_API_Inaccessible
		return public.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return public.Response{
		Status: code,
		Data:   resMsg,
		Msg:    e.GetMsg(code),
	}
}

// SetChatWindow set connChatConn[需要校验删除用户信息]
func (c *ChatService) SetChatWindow() public.Response {
	code := e.SUCCESS
	conn, err := global.ChatConnManager.Get(c.ChatReq.ConnId)
	if err != nil {
		code = e.ChatGPT_Manager_GetConnFail
		return public.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	conn.SetProperty("noyet", "noyet")
	return public.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// PingChatMsg 打印输出信息
func PingChatMsg(msg []openai.ChatCompletionChoice) {
	for _, val := range msg {
		fmt.Printf("+v", val)
	}
}
