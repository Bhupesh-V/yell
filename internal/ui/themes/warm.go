package themes

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var (
	warmBg          = color.NRGBA{R: 222, G: 198, B: 149, A: 255} // Warm Sand
	warmTitle       = color.NRGBA{R: 42, G: 31, B: 26, A: 255}    // Deep Espresso / Roasted Coffee
	warmSubtleText  = color.NRGBA{R: 105, G: 82, B: 71, A: 255}   // Muted Mocha
	warmHoverCircle = color.NRGBA{R: 42, G: 31, B: 26, A: 35}     // Translucent Espresso hover
)

type WarmTheme struct{}

func (t WarmTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return warmBg

	case theme.ColorNameForeground, theme.ColorNamePrimary:
		return warmTitle

	case ColorNameSubtleText:
		return warmSubtleText

	case ColorNameHoverCircle:
		return warmHoverCircle

	default:
		return theme.DefaultTheme().Color(name, theme.VariantLight)
	}
}

func (t WarmTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (t WarmTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t WarmTheme) Size(name fyne.ThemeSizeName) float32 {
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
