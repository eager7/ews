package ws

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"strings"
	"time"
)

var ErrClosed = errors.New("network connection closed")

type MessageHandler func(conn *websocket.Conn, msgType int, MsgContent string, err error)

/*
** websocket读数据线程，数据读取后会调用回调处理，回调函数是同步的，需要用户自己在回调中做异步处理
*/
func readRoutine(conn *websocket.Conn, callback MessageHandler) {
	for {
		typ, body, err := conn.ReadMessage()
		if err != nil {
			logger.Println("read message err, closed:", err)
			callback(conn, typ, string(body), err)
			if err := conn.Close(); err != nil {
				logger.Println("close conn err:", err)
				callback(conn, typ, string(body), err)
			}
			return
		} else {
			logger.Println("message:", typ, string(body))
			callback(conn, typ, string(body), nil)
		}
	}
}

/*
** 写数据到websocket对端
*/
func Write(conn *websocket.Conn, msgType int, message []byte) error {
	_ = conn.SetWriteDeadline(time.Now().Add(time.Second * 5))
	if err := conn.WriteMessage(msgType, message); err != nil {
		logger.Println(err)
		if strings.Contains(err.Error(), "use of closed network connection") {
			return ErrClosed
		}
		return fmt.Errorf("write message err:%v", err)
	}
	return nil
}
