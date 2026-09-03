//go:build linux

package nativewindow

import (
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver"
	"github.com/ebitengine/purego"
)

var (
	x11Once    sync.Once
	x11Ready   bool
	display    uintptr
	rootWindow uintptr

	xMoveWindow           func(display, w uintptr, x, y int32) int32
	xFlush                func(display uintptr) int32
	xTranslateCoordinates func(display, srcW, destW uintptr, srcX, srcY int32, destX, destY *int32, child *uintptr) int32
)

// loadX11 dlopens libX11 directly rather than requiring cgo. It silently
// leaves x11Ready false on a pure-Wayland session with no X11 available,
// so Move becomes a no-op there.
func loadX11() {
	x11Once.Do(func() {
		lib, err := purego.Dlopen("libX11.so.6", purego.RTLD_NOW|purego.RTLD_GLOBAL)
		if err != nil {
			return
		}

		var xOpenDisplay func(name uintptr) uintptr
		var xDefaultRootWindow func(display uintptr) uintptr
		purego.RegisterLibFunc(&xOpenDisplay, lib, "XOpenDisplay")
		purego.RegisterLibFunc(&xDefaultRootWindow, lib, "XDefaultRootWindow")
		purego.RegisterLibFunc(&xMoveWindow, lib, "XMoveWindow")
		purego.RegisterLibFunc(&xFlush, lib, "XFlush")
		purego.RegisterLibFunc(&xTranslateCoordinates, lib, "XTranslateCoordinates")

		display = xOpenDisplay(0)
		if display == 0 {
			return
		}
		rootWindow = xDefaultRootWindow(display)
		x11Ready = true
	})
}

// BeginDrag is unused on X11; positioning is driven by Move.
func BeginDrag(win fyne.Window) bool { return false }

var (
	posMu  sync.Mutex
	posMap = map[uintptr][2]int32{}
)

// Move shifts the window by (dx, dy). The absolute position is queried once
// per window (via XTranslateCoordinates against the root window) and then
// tracked locally, since some window managers reparent windows in ways that
// make repeated position queries mid-drag unreliable.
//
// Note: this moves the client window directly rather than negotiating with
// the window manager (e.g. via _NET_WM_MOVERESIZE), so it may not track
// correctly under window managers that add non-zero frame decoration to
// undecorated windows. It works for the common case of borderless popups.
func Move(win fyne.Window, dx, dy float32) bool {
	nw, ok := win.(driver.NativeWindow)
	if !ok {
		return false
	}

	moved := false
	nw.RunNative(func(ctx any) {
		xc, ok := ctx.(driver.X11WindowContext)
		if !ok || xc.WindowHandle == 0 {
			return
		}
		loadX11()
		if !x11Ready {
			return
		}

		posMu.Lock()
		defer posMu.Unlock()

		pos, seen := posMap[xc.WindowHandle]
		if !seen {
			var absX, absY int32
			var child uintptr
			xTranslateCoordinates(display, xc.WindowHandle, rootWindow, 0, 0, &absX, &absY, &child)
			pos = [2]int32{absX, absY}
		}

		pos[0] += int32(dx)
		pos[1] += int32(dy)
		posMap[xc.WindowHandle] = pos

		xMoveWindow(display, xc.WindowHandle, pos[0], pos[1])
		xFlush(display)
		moved = true
	})
	return moved
}
