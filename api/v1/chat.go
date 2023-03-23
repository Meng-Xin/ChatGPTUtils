package v1

import (
	"chatGPT/service"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func CreateChat(c *gin.Context) {
	// 拿到对话信息创建对话
	var chatInfo *service.ChatService
	err := c.ShouldBindJSON(&chatInfo)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(400, "数据格式错误")
	} else {
		res := chatInfo.AddChatWindow()
		c.JSON(200, res)
	}
}

func GetChat(c *gin.Context) {
	// 拿到对话信息创建对话
	var chatInfo *service.ChatService
	err := c.ShouldBindJSON(&chatInfo)
	if err != nil {
		log.Error(err.Error())
		c.JSON(400, "数据格式错误")
	} else {
		res := chatInfo.GetChatWindow()
		c.JSON(200, res)
	}
}

func GetChatToStream(c *gin.Context) {
	// 拿到对话信息创建对话
	var chatInfo *service.ChatService
	err := c.ShouldBindJSON(&chatInfo)
	if err != nil {
		log.Error(err.Error())
		c.JSON(400, "数据格式错误")
	} else {
		res := chatInfo.GetChatToStream()
		c.JSON(200, res)
	}
}

func SetChat(c *gin.Context) {
	// 拿到对话信息创建对话
	var chatInfo *service.ChatService
	err := c.ShouldBindJSON(&chatInfo)
	if err != nil {
		log.Error(err.Error())
		c.JSON(400, "数据格式错误")
	} else {
		res := chatInfo.SetChatWindow()
		c.JSON(200, res)
	}
}

func DeleteChat(c *gin.Context) {
	// 拿到对话信息创建对话
	var chatInfo *service.ChatService
	err := c.ShouldBindJSON(&chatInfo)
	if err != nil {
		log.Error(err.Error())
		c.JSON(400, "数据格式错误")
	} else {
		res := chatInfo.RemoveChatWindow()
		c.JSON(200, res)
	}
}
