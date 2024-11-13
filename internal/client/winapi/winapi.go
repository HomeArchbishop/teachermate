package winapi

import "syscall"

var (
	user32       = syscall.NewLazyDLL("user32.dll")
	showWindow   = user32.NewProc("ShowWindow")
	findWindow   = user32.NewProc("FindWindowW")
	setWindowPos = user32.NewProc("SetWindowPos")
)
