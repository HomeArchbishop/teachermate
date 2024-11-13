package winapi

import (
	"syscall"
	"unsafe"
)

func FindWindow(className, windowName string) (hwnd syscall.Handle, err error) {
	var cname, wname *uint16
	if className != "" {
		cname, err = syscall.UTF16PtrFromString(className)
		if err != nil {
			return 0, err
		}
	}
	if windowName != "" {
		wname, err = syscall.UTF16PtrFromString(windowName)
		if err != nil {
			return 0, err
		}
	}
	r1, _, e1 := syscall.SyscallN(findWindow.Addr(), uintptr(unsafe.Pointer(cname)), uintptr(unsafe.Pointer(wname)))
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	hwnd = syscall.Handle(r1)
	return
}
