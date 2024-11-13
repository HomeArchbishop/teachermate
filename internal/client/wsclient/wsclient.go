package wsclient

import (
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func StartClient(serverHost, lessonId, studentId string) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	query := url.Values{}
	query.Set("lesson_id", lessonId)
	query.Set("student_id", studentId)
	u := url.URL{
		Scheme:   "ws",
		Host:     serverHost,
		Path:     "/ws",
		RawQuery: query.Encode(),
	}

	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("err: dial: ", err)
	}
	defer c.Close()
	log.Printf("connected to %s", u.String())

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("err read:", err)
				return
			}
			handleMessage(message)
		}
	}()

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("interrupt")
			// cleanly close the connection
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("err: write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
