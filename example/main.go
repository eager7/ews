package main

import (
	"fmt"
	"github.com/eager7/ews/ws"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

func main() {
	internal := gin.New()
	internal.Use(gin.Logger(), gin.Recovery())
	iv1Router := internal.Group("v1")

	iv1Router.GET("/ws", WsHandler)

	server := &http.Server{Addr: "0.0.0.0:2333", Handler: internal}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("listen internal err:", err)
	}
}

func WsHandler(ctx *gin.Context) {
	fmt.Println("client:", ctx.ClientIP())
	err := ws.NewHttpConnection(ctx.Writer, ctx.Request, nil, HandleMessage)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusOK, "hello world!")
	}
	fmt.Println("success connect:", ctx.Request.RemoteAddr)
}

func HandleMessage(conn *websocket.Conn, msgType int, MsgContent string, err error) {
	fmt.Println("receive message:", msgType, MsgContent, err)
	if err != nil {
		return
	}
	_ = ws.Write(conn, 2, []byte(MsgContent))
}
