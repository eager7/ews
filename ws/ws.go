package ws

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"time"
)

var logger = log.New(os.Stdout, "[ws]", log.LstdFlags)

/*
** 将http请求升级为web socket长连接
 */
func NewHttpConnection(w http.ResponseWriter, r *http.Request, respHeader http.Header, callback Handler) error {
	upGrader := websocket.Upgrader{
		HandshakeTimeout: 5 * time.Second,
		ReadBufferSize:   4096,
		WriteBufferSize:  4096,
		CheckOrigin: func(r *http.Request) bool { //跨域检测，前面中间件做了检测的话，这里直接通过即可
			return true
		},
		EnableCompression: true, //压缩
	}

	wsConn, err := upGrader.Upgrade(w, r, respHeader)
	if err != nil {
		logger.Println("upgrade http connect to ws err:", err)
		return err
	}
	go readRoutine(wsConn, callback)
	return err
}
