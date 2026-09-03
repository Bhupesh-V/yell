package components

import (
	"yell/internal/ui/nativewindow"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// DraggableContainer wraps content so that dragging it moves the containing
// window. This is needed for undecorated/splash windows, which have no
// title bar for the OS to handle dragging with.
type DraggableContainer struct {
	widget.BaseWidget
	Content fyne.CanvasObject
	window  fyne.Window

	// wmDrag is true when the window manager took over positioning for the
	// current gesture (see nativewindow.BeginDrag), so Dragged should not
	// also nudge the window itself.
	wmDrag bool
}

func NewDraggableContainer(content fyne.CanvasObject, window fyne.Window) *DraggableContainer {
	d := &DraggableContainer{
		Content: content,
		window:  window,
	}
	d.ExtendBaseWidget(d)
	return d
}

var _ fyne.Draggable = (*DraggableContainer)(nil)
var _ desktop.Mouseable = (*DraggableContainer)(nil)

func (d *DraggableContainer) MouseDown(e *desktop.MouseEvent) {
	if d.window == nil || e.Button != desktop.MouseButtonPrimary {
		return
	}
	d.wmDrag = nativewindow.BeginDrag(d.window)
}

func (d *DraggableContainer) MouseUp(*desktop.MouseEvent) {
	d.wmDrag = false
}

func (d *DraggableContainer) Dragged(e *fyne.DragEvent) {
	if d.window == nil || d.wmDrag {
		return
	}
	nativewindow.Move(d.window, e.Dragged.DX, e.Dragged.DY)
}

func (d *DraggableContainer) DragEnd() {
	d.wmDrag = false
}

func (d *DraggableContainer) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(d.Content)
}
