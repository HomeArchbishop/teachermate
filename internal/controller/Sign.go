package controller

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/homearchbishop/teachermate-auto/internal/service"
)

func signHandler(w http.ResponseWriter, r *http.Request) {
	lessonId := r.URL.Query().Get("lesson_id")
	attendanceId := r.URL.Query().Get("attendance")

	connIdList, webErr := service.GetSubscriptionConns(lessonId)
	if webErr != nil {
		http.Error(w, webErr.Message, webErr.Code)
		return
	}

	connList.mux.Lock()
	for _, connId := range connIdList {
		if conn, ok := connList.conn[connId]; ok {
			// no need to check error,
			// Because PingMessage will check the connection status
			// and close the connection if it is not available.
			conn.WriteMessage(websocket.TextMessage, []byte(attendanceId))
		}
	}
	connList.mux.Unlock()

	w.Write([]byte("ok"))
}
