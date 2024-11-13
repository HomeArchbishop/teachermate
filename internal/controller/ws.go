package controller

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/homearchbishop/teachermate-auto/internal/service"
)

type connListType struct {
	conn map[string]*websocket.Conn
	mux  sync.Mutex
}

var connList = &connListType{conn: make(map[string]*websocket.Conn, 0)}

func wsConnectHandler(w http.ResponseWriter, r *http.Request, conn *websocket.Conn, connId string) {
	connList.mux.Lock()
	connList.conn[connId] = conn
	connList.mux.Unlock()

	lessonId := r.URL.Query().Get("lesson_id")
	studentId := r.URL.Query().Get("student_id")

	webErr := service.SubscribeSignSignal(lessonId, studentId, connId)
	if webErr != nil {
		http.Error(w, webErr.Message, webErr.Code)
		conn.Close()
		return
	}
}

func wsDisconnectHandler(r *http.Request, connId string) {
	connList.mux.Lock()
	delete(connList.conn, connId)
	connList.mux.Unlock()

	lessonId := r.URL.Query().Get("lesson_id")
	studentId := r.URL.Query().Get("student_id")

	// no need to check error
	service.CancelSubscription(lessonId, studentId)
}
