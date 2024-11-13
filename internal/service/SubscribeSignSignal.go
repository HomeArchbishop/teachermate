package service

import (
	"log"

	"github.com/homearchbishop/teachermate-auto/internal/model"
)

func SubscribeSignSignal(lessonId, studentId, connId string) *WebError {
	if lessonId == "" || studentId == "" {
		return NewWebError(400, "empty arguments")
	}

	if err := model.AddSubscription(lessonId, studentId, connId); err != nil {
		log.Println(err)
		return NewWebError(500, "internal error")
	}

	return nil
}
