package ws

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"strings"
	"time"
)

var ErrClosed = errors.New("network connection closed")

type CloseHandler func(code int, text string) error
type MessageHandler func(conn *websocket.Conn, msgType int, MsgContent string)

func readRoutine(conn *websocket.Conn, callback MessageHandler) {
	for {
		typ, body, err := conn.ReadMessage()
		if err != nil {
			logger.Println("read message err, closed:", err)
			if err := conn.Close(); err != nil {
				logger.Println("close conn err:", err)
			}
			return
		} else {
			logger.Println("message:", typ, string(body))
			callback(conn, typ, string(body))
		}
	}
}

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
