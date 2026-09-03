//go:build windows

package nativewindow

import (
	"syscall"
	"unsafe"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver"
)

type winRect struct{ Left, Top, Right, Bottom int32 }

var (
	user32            = syscall.NewLazyDLL("user32.dll")
	procGetWindowRect = user32.NewProc("GetWindowRect")
	procSetWindowPos  = user32.NewProc("SetWindowPos")
)

const (
	swpNoSize     = 0x0001
	swpNoZOrder   = 0x0004
	swpNoActivate = 0x0010
)

// BeginDrag is unused on Windows; positioning is driven by Move.
func BeginDrag(win fyne.Window) bool { return false }

// Move shifts the window by (dx, dy), given in screen pixels.
func Move(win fyne.Window, dx, dy float32) bool {
	nw, ok := win.(driver.NativeWindow)
	if !ok {
		return false
	}

	moved := false
	nw.RunNative(func(ctx any) {
		wc, ok := ctx.(driver.WindowsWindowContext)
		if !ok || wc.HWND == 0 {
			return
		}

		var r winRect
		ret, _, _ := procGetWindowRect.Call(wc.HWND, uintptr(unsafe.Pointer(&r)))
		if ret == 0 {
			return
		}

		newX := r.Left + int32(dx)
		newY := r.Top + int32(dy)
		procSetWindowPos.Call(
			wc.HWND, 0,
			uintptr(newX), uintptr(newY), 0, 0,
			uintptr(swpNoSize|swpNoZOrder|swpNoActivate),
		)
		moved = true
	})
	return moved
}
