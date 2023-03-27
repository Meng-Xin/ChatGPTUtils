package router

import (
	api "chatGPT/api/v1/chat"
	"chatGPT/middle"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	// swagger
	publicGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// V1管理
	v1 := r.Group("/api/v1")
	{
		openaiPri := r.Group("/openai")
		openaiPri.Use()
		v1.GET("/openai/getScenes", api.GetScenes)
		v1.POST("/openai/addScenes", api.AddScenes)
		v1.POST("/openai/saveScenes", api.SaveScenes)
		v1.POST("/openai/scenesChat", api.ScenesChat)
		v1.PUT("/openai/setScenes", api.SetScenes)
		v1.DELETE("/openai/deleteScenes", api.DeleteScenes)
	}
	return r
}
