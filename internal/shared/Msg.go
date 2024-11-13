package shared

import "errors"

type BasicMsgType struct {
	Type string `json:"type"`
}

type SignSignalMsgType struct {
	BasicMsgType
	SignURL string `json:"sign_url"`
}

func CreateSignSignalMsg(signUrl string) (*SignSignalMsgType, error) {
	if signUrl == "" {
		return nil, errors.New("empty attendance query")
	}
	return &SignSignalMsgType{
		BasicMsgType: BasicMsgType{Type: "sign_signal"},
		SignURL:      signUrl,
	}, nil
}
