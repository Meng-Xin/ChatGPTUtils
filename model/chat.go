package model

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	ConnId     uint32 `json:"conn_id"`     //会话id,有就复用conn
	ChatModel  int    `json:"chat_model"`  //会话模型
	HistoryMsg string `json:"history_msg"` //历史聊天记录

	UserId uint `json:"user_id"`
}

func SaveChatScenes(connId uint32, chatModel int, History string) *Chat {
	return &Chat{
		ConnId:     connId,
		ChatModel:  chatModel,
		HistoryMsg: History,
	}
}
