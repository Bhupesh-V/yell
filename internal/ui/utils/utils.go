package utils

import (
	"image/color"
	"path"
	"yell/internal/assets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

var InterFont = fyne.NewStaticResource("font.ttf", assets.FontBytes)

func Spacer(w, h float32) fyne.CanvasObject {
	rect := canvas.NewRectangle(color.Transparent)
	rect.SetMinSize(fyne.NewSize(w, h))
	return rect
}

func GetEmbeddedEmojiImage(emoji string, targetSize float32) fyne.CanvasObject {
	filePath := assets.GetEmojiFilePath(emoji)

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

// Helper function to map string to Fyne's ThemeVariant
// func ParseThemeMode(modeStr string) fyne.ThemeVariant {
// 	switch strings.ToLower(strings.TrimSpace(modeStr)) {
// 	case "light":
// 		return theme.VariantLight
// 	case "dark":
// 		return theme.VariantDark
// 	case "cyberpunk":
// 		return components.CyberpunkTheme
// 	default:
// 		// Default mode
// 		return theme.VariantDark
// 	}
// }

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
