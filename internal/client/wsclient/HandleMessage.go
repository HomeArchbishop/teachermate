package wsclient

import (
	"encoding/json"
	"log"

	"github.com/homearchbishop/teachermate-auto/internal/shared"
)

func handleMessage(msg []byte) {
	log.Println(string(msg))
	// parse msg
	basicMsg, err := parseMessage[*shared.BasicMsgType](msg)
	if err != nil {
		log.Println(err)
		return
	}
	switch basicMsg.Type {
	case "sign_signal":
		handleSignSignalMessage(msg)
	}
}

func parseMessage[T shared.Msg](msg []byte) (T, error) {
	var signSignalMsg T
	if err := json.Unmarshal(msg, &signSignalMsg); err != nil {
		log.Println(err)
		return signSignalMsg, err
	}
	return signSignalMsg, nil
}

func handleSignSignalMessage(msg []byte) {
	signSignalMsg, err := parseMessage[*shared.SignSignalMsgType](msg)
	if err != nil {
		log.Printf("failed to parse sign_signal msg: %s\n", string(msg))
		return
	}
	if handler, ok := registedHandler["sign_signal"]; ok {
		handler(signSignalMsg)
	} else {
		log.Println("no handler for sign_signal")
	}
}

var registedHandler = make(map[string]func(shared.Msg))

func RegisterHandler[T shared.Msg](typeName string, handler func(msg T)) {
	registedHandler[typeName] = func(msg shared.Msg) {
		handler(msg.(T))
	}
}
