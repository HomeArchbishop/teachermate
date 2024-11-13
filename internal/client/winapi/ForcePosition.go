package winapi

import "syscall"

const (
	SWP_SHOWWINDOW = 0x0040
	HWND_TOPMOST   = ^uintptr(0) // -1

	SW_RESTORE = 9
)

func ForcePosition(hwnd syscall.Handle, x, y, width, height uintptr) error {
	syscall.SyscallN(showWindow.Addr(), uintptr(hwnd), SW_RESTORE)
	syscall.SyscallN(setWindowPos.Addr(), uintptr(hwnd), HWND_TOPMOST, x, y, width, height, SWP_SHOWWINDOW)
	return nil
}
