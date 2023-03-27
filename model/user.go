package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     string `json:"uuid"`      //UUID
	UserName string `json:"user_name"` //账号
	NickName string `json:"nick_name"` //昵称
	Password string `json:"password"`  //密码

	ChatConfig []Chat `json:"chat_msg"`
}
