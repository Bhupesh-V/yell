package themes

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type CyberpunkTheme struct{}

func (t CyberpunkTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.NRGBA{R: 9, G: 2, B: 33, A: 255} // Deep neon purple

	case theme.ColorNameForeground, theme.ColorNamePrimary:
		return color.NRGBA{R: 0, G: 255, B: 204, A: 255} // Neon Cyan (Title)

	// Explicitly handle the custom keys used by the message and close button
	case ColorNameSubtleText:
		return color.NRGBA{R: 255, G: 105, B: 180, A: 255} // Hot Pink (Message & '✕' icon)

	case ColorNameHoverCircle:
		return color.NRGBA{R: 0, G: 255, B: 204, A: 40} // Translucent Cyan (Close hover state)

	default:
		// Always fall back to the default Fyne theme for unmapped standard colors
		return theme.DefaultTheme().Color(name, variant)
	}
}

func (t CyberpunkTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (t CyberpunkTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t CyberpunkTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNamePadding:
		return 12
	case theme.SizeNameInnerPadding:
		return 2 // Reduces internal widget padding/bounding box height
	case theme.SizeNameText:
		return 14 // Matches standard text scaling
	default:
		return theme.DefaultTheme().Size(name)
	}
}
