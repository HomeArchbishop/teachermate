package service

type WebError struct {
	Code    int
	Message string
}

func NewWebError(code int, msg string) *WebError {
	return &WebError{Code: code, Message: msg}
}
