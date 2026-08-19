package components

import (
	"image/color"
	"yell/internal/ui/themes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// --- Custom Circular Hover Close Button ---
type circularCloseButton struct {
	widget.BaseWidget
	onTapped func()
	hovered  bool
	circle   *canvas.Circle
	text     *canvas.Text
}

func NewCircularCloseButton(onTapped func()) *circularCloseButton {
	b := &circularCloseButton{
		onTapped: onTapped,
	}
	b.ExtendBaseWidget(b)
	return b
}

func (b *circularCloseButton) MouseIn(*desktop.MouseEvent) {
	b.hovered = true
	b.Refresh()
}

func (b *circularCloseButton) MouseOut() {
	b.hovered = false
	b.Refresh()
}

func (b *circularCloseButton) MouseMoved(*desktop.MouseEvent) {}

func (b *circularCloseButton) Tapped(*fyne.PointEvent) {
	if b.onTapped != nil {
		b.onTapped()
	}
}

func (b *circularCloseButton) CreateRenderer() fyne.WidgetRenderer {
	th := theme.Current()
	variant := fyne.CurrentApp().Settings().ThemeVariant()

	b.circle = canvas.NewCircle(color.Transparent)
	b.text = canvas.NewText("✕", th.Color(themes.ColorNameSubtleText, variant))
	b.text.TextSize = 12
	b.text.Alignment = fyne.TextAlignCenter

	return &circularCloseButtonRenderer{
		btn:     b,
		objects: []fyne.CanvasObject{b.circle, b.text},
	}
}

type circularCloseButtonRenderer struct {
	btn     *circularCloseButton
	objects []fyne.CanvasObject
}

func (r *circularCloseButtonRenderer) Destroy() {}

func (r *circularCloseButtonRenderer) Layout(size fyne.Size) {
	minDim := size.Width
	if size.Height < minDim {
		minDim = size.Height
	}

	posX := (size.Width - minDim) / 2
	posY := (size.Height - minDim) / 2

	r.btn.circle.Move(fyne.NewPos(posX, posY))
	r.btn.circle.Resize(fyne.NewSize(minDim, minDim))

	txtMin := r.btn.text.MinSize()
	textX := (size.Width - txtMin.Width) / 2
	textY := (size.Height - txtMin.Height) / 2

	r.btn.text.Move(fyne.NewPos(textX, textY))
	r.btn.text.Resize(txtMin)
}

func (r *circularCloseButtonRenderer) MinSize() fyne.Size {
	return fyne.NewSize(28, 28)
}

func (r *circularCloseButtonRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

// Refresh fetches colors directly from the registered global theme
func (r *circularCloseButtonRenderer) Refresh() {
	th := theme.Current()
	variant := fyne.CurrentApp().Settings().ThemeVariant()

	if r.btn.hovered {
		r.btn.circle.FillColor = th.Color(themes.ColorNameHoverCircle, variant)
		r.btn.text.Color = th.Color(theme.ColorNameForeground, variant)
	} else {
		r.btn.circle.FillColor = color.Transparent
		r.btn.text.Color = th.Color(themes.ColorNameSubtleText, variant)
	}
	r.btn.circle.Refresh()
	r.btn.text.Refresh()
}
