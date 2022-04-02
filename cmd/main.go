package main

import (
	"chat_room/router"
	gin2 "github.com/gin-gonic/gin"
)

func main() {
	g := gin2.Default()
	router.Init(g)
	g.Run("localhost:8080")
}