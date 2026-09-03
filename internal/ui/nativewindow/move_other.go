//go:build !windows && !darwin && !linux

package nativewindow

import "fyne.io/fyne/v2"

// BeginDrag and Move are no-ops on platforms without a native move
// implementation (e.g. mobile, wasm).
func BeginDrag(win fyne.Window) bool            { return false }
func Move(win fyne.Window, dx, dy float32) bool { return false }
