package themes

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// Palette defines the custom color set for a theme
type Palette struct {
	Background  color.Color
	Foreground  color.Color
	SubtleText  color.Color
	HoverCircle color.Color
	Variant     fyne.ThemeVariant
}

type BaseTheme struct {
	Palette Palette
}

func NewBaseTheme(p Palette) *BaseTheme {
	return &BaseTheme{Palette: p}
}

func (t *BaseTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return t.Palette.Background
	case theme.ColorNameForeground, theme.ColorNamePrimary:
		return t.Palette.Foreground
	case ColorNameSubtleText:
		return t.Palette.SubtleText
	case ColorNameHoverCircle:
		return t.Palette.HoverCircle
	default:
		return theme.DefaultTheme().Color(name, t.Palette.Variant)
	}
}

func (t *BaseTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (t *BaseTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t *BaseTheme) Size(name fyne.ThemeSizeName) float32 {
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
