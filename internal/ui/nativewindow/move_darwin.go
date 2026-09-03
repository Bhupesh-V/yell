//go:build darwin

package nativewindow

import (
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver"
	"github.com/ebitengine/purego"
	"github.com/ebitengine/purego/objc"
)

type nsPoint struct{ X, Y float64 }
type nsSize struct{ W, H float64 }
type nsRect struct {
	Origin nsPoint
	Size   nsSize
}

var (
	cocoaOnce    sync.Once
	frameSel     objc.SEL
	setOriginSel objc.SEL
)

func loadCocoa() {
	cocoaOnce.Do(func() {
		if _, err := purego.Dlopen("/System/Library/Frameworks/Cocoa.framework/Cocoa", purego.RTLD_GLOBAL|purego.RTLD_LAZY); err != nil {
			return
		}
		frameSel = objc.RegisterName("frame")
		setOriginSel = objc.RegisterName("setFrameOrigin:")
	})
}

// BeginDrag is unused on macOS; NSWindow positioning is driven by Move.
func BeginDrag(win fyne.Window) bool { return false }

// Move shifts the window by (dx, dy), given in Fyne's top-left-origin,
// downward-positive coordinate space. AppKit's origin is bottom-left, so dy
// is inverted when computing the new frame origin.
func Move(win fyne.Window, dx, dy float32) bool {
	nw, ok := win.(driver.NativeWindow)
	if !ok {
		return false
	}

	moved := false
	nw.RunNative(func(ctx any) {
		mc, ok := ctx.(driver.MacWindowContext)
		if !ok || mc.NSWindow == 0 {
			return
		}
		loadCocoa()
		if frameSel == 0 {
			return
		}

		nsWindow := objc.ID(mc.NSWindow)
		frame := objc.Send[nsRect](nsWindow, frameSel)
		newOrigin := nsPoint{
			X: frame.Origin.X + float64(dx),
			Y: frame.Origin.Y - float64(dy),
		}
		nsWindow.Send(setOriginSel, newOrigin)
		moved = true
	})
	return moved
}
