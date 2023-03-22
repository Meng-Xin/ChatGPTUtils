package router

import (
	api "chatGPT/api/v1"
	"chatGPT/middle"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middle.LoggerToFile())

	//公共路由组
	publicGroup := r.Group("")
	publicGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "Success")
		log.Info("Test")
	})
	publicGroup.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello welcome Meng-Xin")
	})
	// V1管理
	v1 := r.Group("/api/v1")
	{
		v1.POST("/chat/addWindow", api.CreateChat)
		v1.GET("/chat/getWindow", api.GetChat)
		v1.GET("/chat/getStream", api.GetChatToStream)
		v1.PUT("/chat/setWindow", api.SetChat)
		v1.DELETE("/chat/deleteWindow", api.DeleteChat)
	}
	return r
}
