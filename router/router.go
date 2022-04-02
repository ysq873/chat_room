package router

import (
	"chat_room/api/home"
	"chat_room/api/user"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	// 读取该文件下
	router.LoadHTMLGlob("view/**/*")

	// 用户路由组
	userRouterGroup := router.Group("/user")
	// 使用中间件
	// 1、 加密
	// 2、 鉴权
	//router.Use()
	{
		userRouterGroup.GET("/list", user.OnLineList)
		userRouterGroup.POST("/register", user.Register)
		userRouterGroup.POST("/login", user.Login)
		userRouterGroup.POST("/sendMessage", user.SendMessage)
		userRouterGroup.POST("/sendMessageAll", user.SendMessageAll)
	}
	// 网站路由组
	homeRouterGroup := router.Group("/room")
	{
		homeRouterGroup.POST("/index", home.Index)
	}

}