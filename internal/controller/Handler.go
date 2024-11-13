package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	connId := uuid.New().String()

	conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	go func() {
		for {
			// ping
			time.Sleep(30 * time.Second)
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				conn.Close()
				return
			}
		}
	}()

	conn.SetCloseHandler(func(code int, text string) error {
		wsDisconnectHandler(w, r, connId)
		return nil
	})

	wsConnectHandler(w, r, conn, connId)
}

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	subPath := r.URL.Path[len("/api/"):]
	switch subPath {
	case "sign":
		signHandler(w, r)
	}
}
