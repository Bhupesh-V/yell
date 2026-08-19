package ui

import (
	"fmt"
	"image/color"
	"path"
	"strings"
	"yell/internal/assets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var interFontResource = fyne.NewStaticResource("font.ttf", assets.FontBytes)

type ToastTheme struct{}

func (t ToastTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{R: 24, G: 24, B: 36, A: 255}
	case theme.ColorNameForeground, theme.ColorNamePrimary:
		return color.White
	default:
		// Always fall back to the Dark variant color definitions
		return theme.DefaultTheme().Color(name, theme.VariantDark)
	}
}

// Serves Inter custom font across all operating systems
func (t ToastTheme) Font(style fyne.TextStyle) fyne.Resource {
	if interFontResource != nil && len(interFontResource.StaticContent) > 0 {
		return interFontResource
	}
	return theme.DefaultTheme().Font(style)
}

func (t ToastTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t ToastTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNamePadding:
		return 12
	case theme.SizeNameInnerPadding:
		return 2
	case theme.SizeNameText:
		return 14
	default:
		return theme.DefaultTheme().Size(name)
	}
}

// Custom VBox for exact pixel spacing between items
type TightVBoxLayout struct {
	Spacing float32
}

func (t *TightVBoxLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	y := float32(0)
	for _, child := range objects {
		childSize := child.MinSize()
		child.Resize(childSize)
		child.Move(fyne.NewPos(0, y))
		y += childSize.Height + t.Spacing
	}
}

func (t *TightVBoxLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	var width, height float32
	for i, child := range objects {
		min := child.MinSize()
		if min.Width > width {
			width = min.Width
		}
		height += min.Height
		if i > 0 {
			height += t.Spacing
		}
	}
	return fyne.NewSize(width, height)
}

// --- Custom Circular Hover Close Button ---
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
	b.circle = canvas.NewCircle(color.Transparent)
	b.text = canvas.NewText("✕", color.NRGBA{R: 180, G: 180, B: 180, A: 255})
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

func (r *circularCloseButtonRenderer) Refresh() {
	if r.btn.hovered {
		r.btn.circle.FillColor = color.NRGBA{R: 255, G: 255, B: 255, A: 35}
		r.btn.text.Color = color.White
	} else {
		r.btn.circle.FillColor = color.Transparent
		r.btn.text.Color = color.NRGBA{R: 180, G: 180, B: 180, A: 255}
	}
	r.btn.circle.Refresh()
	r.btn.text.Refresh()
}

// ------------------------------------------

func getEmojiFilePath(emoji string) string {
	var fullHex []string
	var cleanHex []string

	for _, r := range emoji {
		hexVal := fmt.Sprintf("%x", r)
		fullHex = append(fullHex, hexVal)
		if r != 0xfe0f {
			cleanHex = append(cleanHex, hexVal)
		}
	}

	cleanPath := path.Join("assets", "emojis", strings.Join(cleanHex, "-")+".png")
	if _, err := assets.Emojis.ReadFile(cleanPath); err == nil {
		return cleanPath
	}

	fullPath := path.Join("assets", "emojis", strings.Join(fullHex, "-")+".png")
	if _, err := assets.Emojis.ReadFile(fullPath); err == nil {
		return fullPath
	}

	return ""
}

func GetEmbeddedEmojiImage(emoji string, targetSize float32) fyne.CanvasObject {
	filePath := getEmojiFilePath(emoji)

	if filePath != "" {
		data, err := assets.Emojis.ReadFile(filePath)
		if err == nil {
			res := fyne.NewStaticResource(path.Base(filePath), data)
			img := canvas.NewImageFromResource(res)
			img.SetMinSize(fyne.NewSize(targetSize, targetSize))
			img.FillMode = canvas.ImageFillContain
			return img
		}
	}

	fallbackText := canvas.NewText(emoji, color.White)
	fallbackText.TextSize = targetSize
	return fallbackText
}
