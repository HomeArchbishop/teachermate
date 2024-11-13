package shared

import "errors"

type Msg interface {
	GetType() string
}

type BasicMsgType struct {
	Type string `json:"type"`
}

type SignSignalMsgType struct {
	Type    string `json:"type"`
	SignUrl string `json:"sign_url"`
}

func (msg *BasicMsgType) GetType() string {
	return msg.Type
}

func (msg *SignSignalMsgType) GetType() string {
	return msg.Type
}

func CreateSignSignalMsg(signUrl string) (*SignSignalMsgType, error) {
	if signUrl == "" {
		return nil, errors.New("empty attendance query")
	}
	return &SignSignalMsgType{
		Type:    "sign_signal",
		SignUrl: signUrl,
	}, nil
}
