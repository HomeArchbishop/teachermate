package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/homearchbishop/teachermate-auto/internal/service"
	"github.com/homearchbishop/teachermate-auto/internal/shared"
)

func signHandler(w http.ResponseWriter, r *http.Request) {
	lessonId := r.URL.Query().Get("lesson_id")
	signUrl := r.URL.Query().Get("attendance")

	connIdList, webErr := service.GetSubscriptionConns(lessonId)
	if webErr != nil {
		http.Error(w, webErr.Message, webErr.Code)
		return
	}

	signSignalMsg, err := shared.CreateSignSignalMsg(signUrl)
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), 500)
		return
	}

	signSignalMsgJson, err := json.Marshal(signSignalMsg)
	if err != nil {
		http.Error(w, "internal error", 500)
		return
	}

	connList.mux.Lock()
	for _, connId := range connIdList {
		if conn, ok := connList.conn[connId]; ok {
			// no need to check error,
			// Because PingMessage will check the connection status
			// and close the connection if it is not available.
			conn.WriteMessage(websocket.TextMessage, signSignalMsgJson)
		}
	}
	connList.mux.Unlock()

	w.Write([]byte("ok"))
}
