package router

import (
	api "chatGPT/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use()

	//公共路由组
	publicGroup := r.Group("")
	publicGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "Success")
	})
	publicGroup.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello welcome Meng-Xin")
	})
	// V1管理
	v1 := r.Group("/api/v1")
	{
		v1.GET("chat", api.CreateChat)
	}
	return r
}
