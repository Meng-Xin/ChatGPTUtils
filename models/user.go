package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     string `json:"uuid"`      //UUID
	NickName string `json:"nick_name"` //昵称
	Account  string `json:"account"`   //账号
	Password string `json:"password"`  //密码
	ConnId   uint32 `json:"conn_id"`   //会话id,有就复用conn

	ChatConfig []Chat `json:"chat_msg"`
}
