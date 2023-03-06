package v1

import (
	"chatGPT/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateChat(c *gin.Context) {
	// 拿到对话信息创建对话
	var chatInfo *service.ChatReq
	err := c.ShouldBindJSON(&chatInfo)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(400, "数据格式错误")
	} else {
		res := chatInfo.AddChatWindow()
		c.JSON(200, res)
	}
}
