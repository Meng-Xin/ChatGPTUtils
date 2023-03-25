package chat

import (
	"chatGPT/models/request"
	"chatGPT/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary	创建会话场景
// @Produce	json
// @Param		request	body		request.AddToScenesRequest	true	"新建场景：1聊天、2画图可选"
// @Success	200		{object}	public.Response				"成功"
// @Failure	400		{object}	public.Response				"请求错误"
// @Router		/api/v1/openai/addScenes [post]
func AddScenes(c *gin.Context) {
	// 拿到对话信息创建对话
	var addToConn request.AddToScenesRequest
	if err := c.ShouldBind(&addToConn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	chatInfo := &service.ChatService{}
	res := chatInfo.AddToScenes(addToConn)
	c.JSON(http.StatusOK, res)
}

// @Summary	聊天场景
// @Produce	json
// @Param		request	body		request.ChatToScenesRequest	true	"聊天场景"
// @Success	200		{object}	public.Response				"成功"
// @Failure	400		{object}	public.Response				"请求错误"
// @Router		/api/v1/openai/scenesChat [post]
func ScenesChat(c *gin.Context) {
	// 拿到对话信息创建对话
	var chat request.ChatToScenesRequest
	if err := c.ShouldBindJSON(&chat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var server *service.ChatService
	res := server.ChatToScenes(chat)
	c.JSON(http.StatusOK, res)
}

func GetScenes(c *gin.Context) {

}

func SetScenes(c *gin.Context) {

}

func DeleteScenes(c *gin.Context) {

}

//func GetChatToStream(c *gin.Context) {
//	// 拿到对话信息创建对话
//	var chatInfo *service.ChatService
//	if err := c.ShouldBindJSON(&chatInfo); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	res := chatInfo.GetChatToStream()
//	c.JSON(http.StatusOK, res)
//}
//
//func SetChat(c *gin.Context) {
//	// 拿到对话信息创建对话
//	var chatInfo *service.ChatService
//	if err := c.ShouldBindJSON(&chatInfo); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	res := chatInfo.SetChatWindow()
//	c.JSON(http.StatusOK, res)
//}
//
//func DeleteChat(c *gin.Context) {
//	// 拿到对话信息创建对话
//	var chatInfo *service.ChatService
//	if err := c.ShouldBindJSON(&chatInfo); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	res := chatInfo.RemoveChatWindow()
//	c.JSON(http.StatusOK, res)
//}
