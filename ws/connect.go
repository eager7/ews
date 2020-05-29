package ws

import (
	"github.com/gorilla/websocket"
	"time"
)

type Handler func(conn *websocket.Conn, msgType int, MsgContent string)

func readRoutine(conn *websocket.Conn, callback Handler) {
	for {
		select {
		default:
			typ, body, err := conn.ReadMessage()
			if err != nil {
				logger.Println("read message err, closed:", err)
			} else {
				logger.Println("message:", typ, string(body))
				go callback(conn, typ, string(body))
			}
		}
	}
}

func Write(conn *websocket.Conn, msgType int, message []byte) error {
	if err := conn.SetWriteDeadline(time.Now().Add(time.Second * 3)); err != nil {
		logger.Println("set write dead line err:", err)
		return err
	}
	if err := conn.WriteMessage(msgType, message); err != nil {
		logger.Println("write message err:", err)
		return err
	}
	return nil
}
