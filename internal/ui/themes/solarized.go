package themes

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// Official Solarized Palette Constants
var (
	solBase03 = color.NRGBA{R: 0, G: 43, B: 54, A: 255}     // Dark background
	solBase02 = color.NRGBA{R: 7, G: 54, B: 66, A: 255}     // Dark highlights / hover
	solBase01 = color.NRGBA{R: 88, G: 110, B: 117, A: 255}  // Light content secondary text
	solBase00 = color.NRGBA{R: 101, G: 123, B: 131, A: 255} // Dark content secondary text
	solBase0  = color.NRGBA{R: 131, G: 148, B: 150, A: 255} // Dark content primary text
	solBase1  = color.NRGBA{R: 147, G: 161, B: 161, A: 255} // Light content primary text
	solBase2  = color.NRGBA{R: 238, G: 232, B: 213, A: 255} // Light highlights / hover
	solBase3  = color.NRGBA{R: 253, G: 246, B: 227, A: 255} // Light background
	solYellow = color.NRGBA{R: 181, G: 137, B: 0, A: 255}   // Accent
	solCyan   = color.NRGBA{R: 42, G: 161, B: 152, A: 255}  // Accent
)

// --- Solarized Dark Theme ---
type SolarizedDarkTheme struct{}

func (t SolarizedDarkTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return solBase03
	case theme.ColorNameForeground, theme.ColorNamePrimary:
		return solBase0
	case ColorNameSubtleText:
		return solBase00
	case ColorNameHoverCircle:
		return color.NRGBA{R: 7, G: 54, B: 66, A: 180} // Base02 hover tint
	default:
		return theme.DefaultTheme().Color(name, theme.VariantDark)
	}
}

func (t SolarizedDarkTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (t SolarizedDarkTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t SolarizedDarkTheme) Size(name fyne.ThemeSizeName) float32 {
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

// --- Solarized Light Theme ---
type SolarizedLightTheme struct{}

func (t SolarizedLightTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return solBase3
	case theme.ColorNameForeground, theme.ColorNamePrimary:
		return solBase01
	case ColorNameSubtleText:
		return solBase1
	case ColorNameHoverCircle:
		return color.NRGBA{R: 238, G: 232, B: 213, A: 180} // Base2 hover tint
	default:
		return theme.DefaultTheme().Color(name, theme.VariantLight)
	}
}

func (t SolarizedLightTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (t SolarizedLightTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t SolarizedLightTheme) Size(name fyne.ThemeSizeName) float32 {
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
