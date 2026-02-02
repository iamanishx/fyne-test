package ui

import "syscall"

func setWindowsWindowStyle(hwnd uintptr) {
	user32 := syscall.NewLazyDLL("user32.dll")
	getWindowLong := user32.NewProc("GetWindowLongPtrW")
	setWindowLong := user32.NewProc("SetWindowLongPtrW")
	setWindowPos := user32.NewProc("SetWindowPos")
	setLayered := user32.NewProc("SetLayeredWindowAttributes")

	const (
		wsOverlapped = 0x00000000
		wsPopup      = 0x80000000
		wsCaption    = 0x00C00000
		wsThickFrame = 0x00040000
		wsMinimize   = 0x20000000
		wsMaximize   = 0x01000000
		wsSysMenu    = 0x00080000
		wsExLayered  = 0x00080000
		swpFrame     = 0x0020
		swpNoMove    = 0x0002
		swpNoSize    = 0x0001
		swpNoZOrder  = 0x0004
		lwaAlpha     = 0x00000002
	)

	gwlStyle := int32(-16)
	gwlExStyle := int32(-20)

	style, _, _ := getWindowLong.Call(hwnd, uintptr(gwlStyle))
	style &^= wsCaption | wsThickFrame | wsMinimize | wsMaximize | wsSysMenu
	style |= wsOverlapped | wsPopup
	setWindowLong.Call(hwnd, uintptr(gwlStyle), style)

	exStyle, _, _ := getWindowLong.Call(hwnd, uintptr(gwlExStyle))
	exStyle |= wsExLayered
	setWindowLong.Call(hwnd, uintptr(gwlExStyle), exStyle)

	setWindowPos.Call(hwnd, 0, 0, 0, 0, 0, swpFrame|swpNoMove|swpNoSize|swpNoZOrder)
	setLayered.Call(hwnd, 0, uintptr(220), uintptr(lwaAlpha))
}

func setX11WindowHints(_ uintptr) {}

func setWaylandWindowHints(_ uintptr) {}
