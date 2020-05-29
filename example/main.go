package main

import (
	"fmt"
	"github.com/eager7/ews/ws"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	internal := gin.New()
	internal.Use(gin.Logger(), gin.Recovery())
	iv1Router := internal.Group("v1")

	iv1Router.GET("/ws", WsHandler)
	iv1Router.POST("/ws", WsHandler)

	server := &http.Server{Addr: "0.0.0.0:2333", Handler: internal}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("listen internal err:", err)
	}
}

func WsHandler(ctx *gin.Context) {
	fmt.Println("client:", ctx.ClientIP())
	err := ws.NewHttpConnection(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusOK, "hello world!")
	}
	fmt.Println("success connect:", ctx.Request.RemoteAddr)
}
