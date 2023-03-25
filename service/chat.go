package service

import (
	chatNet2 "chatGPT/core/conn"
	"chatGPT/global"
	"chatGPT/models/request"
	"chatGPT/models/response"
	"chatGPT/pkg/e"
	"chatGPT/pkg/public"
	"chatGPT/utils"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	log "github.com/sirupsen/logrus"
)

type ChatService struct {
}

// AddToScenes 新建场景
func (c *ChatService) AddToScenes(req request.AddToScenesRequest) public.Response {
	code := e.SUCCESS
	// 使用ChatConnManager 进行调用
	conn := chatNet2.NewChatConn(global.SourceConnID.GetConnID(), req)
	if conn == nil {
		code = e.ChatGPT_API_Create_Failed
		return public.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// resp 封装消息
	resp := response.AddToScenesResponse{ConnId: conn.GetConnID()}
	return public.Response{
		Status: code,
		Data:   resp,
		Msg:    e.GetMsg(code),
	}
}

// ChatToScenes 聊天场景
func (c *ChatService) ChatToScenes(data request.ChatToScenesRequest) public.Response {
	code := e.SUCCESS
	// 获取已创建的会话连接
	conn, err := global.ChatConnManager.Get(data.ConnId)
	if err != nil {
		code = e.ChatGPT_Manager_GetConnFail
		return public.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 发送消息
	resData, err := conn.SendMsg(data.ChatGPT)
	if err != nil {
		log.Error(err.Error())
		code = e.ChatGPT_API_Inaccessible
		return public.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	msg := utils.GetMsg(resData)
	sce := conn.GetScenes().(*chatNet2.Scenes)
	msg.Name = sce.ChatGPT.Name
	return public.Response{
		Status: code,
		Data:   msg,
		Msg:    e.GetMsg(code),
	}
}

func (c *ChatService) PaintToScenes(data request.PaintToScenesRequest) {

}

// SetChatWindow set connChatConn[需要校验删除用户信息]
//func (c *ChatService) SetChatWindow() public.Response {
//	code := e.SUCCESS
//	conn, err := global.ChatConnManager.Get(c.ConnId)
//	if err != nil {
//		code = e.ChatGPT_Manager_GetConnFail
//		return public.Response{
//			Status: code,
//			Msg:    e.GetMsg(code),
//		}
//	}
//	conn.SetProperty("noyet", "noyet")
//	return public.Response{
//		Status: code,
//		Msg:    e.GetMsg(code),
//	}
//}

// GetChatToStream get the created link and send msg[需要校验删除用户信息]
//func (c *ChatService) GetChatToStream() public.Response {
//	code := e.SUCCESS
//	// 获取已创建的会话连接
//	conn, err := global.ChatConnManager.Get(c.ConnId)
//	if err != nil {
//		code = e.ChatGPT_Manager_GetConnFail
//		return public.Response{
//			Status: code,
//			Msg:    e.GetMsg(code),
//		}
//	}
//	// 发送消息
//	msgChan := make(chan string, 1)
//	go func() {
//		defer func() {
//			if errMsg := recover(); errMsg != nil {
//				log.Error(errMsg)
//			}
//			if err != nil {
//				code = e.ChatGPT_Manager_StreamFail
//			}
//			close(msgChan)
//		}()
//		err = conn.(*chatNet2.ChatConnection).SendMsgToChatStream(c.ChatGPT.Msg)
//	}()
//	return public.Response{
//		Status: code,
//		Data:   <-msgChan,
//		Msg:    e.GetMsg(code),
//	}
//}

// RemoveChatWindow remove chatGPT Windows And remove chatConn[需要校验删除用户信息]
func (c *ChatService) RemoveChatWindow(connId uint32) public.Response {
	code := e.SUCCESS
	// 删除对应conn
	conn, err := global.ChatConnManager.Get(connId)
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

// PingChatMsg 打印输出信息
func PingChatMsg(msg []openai.ChatCompletionChoice) {
	for _, val := range msg {
		fmt.Printf("+%v", val)
	}
}
