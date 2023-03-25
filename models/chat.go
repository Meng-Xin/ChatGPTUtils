package models

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	ChatModel  int    `json:"chat_model"`  //会话模型
	HistoryMsg string `json:"history_msg"` //历史聊天记录

	UserId uint `json:"user_id"`
}
