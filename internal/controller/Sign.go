package controller

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/homearchbishop/teachermate-auto/internal/service"
)

func signHandler(w http.ResponseWriter, r *http.Request) {
	lesson_id := r.URL.Query().Get("lesson_id")
	attendance_id := r.URL.Query().Get("attendance")

	connIdList, webErr := service.GetSubscriptionConns(lesson_id)
	if webErr != nil {
		http.Error(w, webErr.Message, webErr.Code)
		return
	}

	connList.mux.Lock()
	for _, connId := range connIdList {
		if conn, ok := connList.conn[connId]; ok {
			conn.WriteMessage(websocket.TextMessage, []byte(attendance_id))
		}
	}
	connList.mux.Unlock()

	w.Write([]byte("ok"))
}
