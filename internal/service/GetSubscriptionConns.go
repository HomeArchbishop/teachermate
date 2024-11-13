package service

import (
	"log"

	"github.com/homearchbishop/teachermate-auto/internal/model"
)

func GetSubscriptionConns(lessonId string) ([]string, *WebError) {
	if lessonId == "" {
		return nil, NewWebError(400, "empty arguments")
	}

	conns, err := model.GetSubscription4Lesson(lessonId)
	if err != nil {
		log.Println(err)
		return nil, NewWebError(500, "internal error")
	}

	return conns, nil
}
