package service

import (
	"log"

	"github.com/homearchbishop/teachermate-auto/internal/model"
)

func CancelSubscription(lessonId, studentId string) *WebError {
	if lessonId == "" || studentId == "" {
		return NewWebError(400, "empty arguments")
	}

	if err := model.RemoveSubscription(lessonId, studentId); err != nil {
		log.Println(err)
		return NewWebError(500, "internal error")
	}

	return nil
}
