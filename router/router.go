package router

import (
	"chat_room/controller/home"
	"chat_room/controller/user"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	// 读取该文件下
	router.LoadHTMLGlob("view/**/*")

	// 用户路由组
	userRouterGroup := router.Group("/user")
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