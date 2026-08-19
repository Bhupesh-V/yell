package components

import (
	"image/color"
	"yell/internal/ui/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

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
	if utils.InterFont != nil && len(utils.InterFont.StaticContent) > 0 {
		return utils.InterFont
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
