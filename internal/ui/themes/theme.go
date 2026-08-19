package themes

import (
	"image/color"
	"yell/internal/ui/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// Custom semantic color tokens
const (
	ColorNameSubtleText  fyne.ThemeColorName = "yell:color:subtleText"
	ColorNameHoverCircle fyne.ThemeColorName = "yell:color:hoverCircle"
)

// ThemeType represents your own custom theme identifiers
type ThemeType string

const (
	ThemeDark           ThemeType = "dark"
	ThemeLight          ThemeType = "light"
	ThemeCyberpunk      ThemeType = "cyberpunk"
	ThemeSolarizedDark  ThemeType = "solarized-dark"
	ThemeSolarizedLight ThemeType = "solarized-light"
	ThemeWarm           ThemeType = "warm"
)

func ParseTheme(input string) ThemeType {
	switch input {
	case "light":
		return ThemeLight
	case "cyberpunk":
		return ThemeCyberpunk
	case "solarized-dark":
		return ThemeSolarizedDark
	case "solarized-light":
		return ThemeSolarizedLight
	case "warm":
		return ThemeWarm
	default:
		return ThemeDark
	}
}

// GetTheme returns the actual fyne.Theme implementation based on ThemeType
func GetTheme(t ThemeType) fyne.Theme {
	switch t {
	case ThemeLight:
		return NewAppTheme(theme.VariantLight)
	case ThemeCyberpunk:
		return &CyberpunkTheme{}
	case ThemeSolarizedDark:
		return &SolarizedDarkTheme{}
	case ThemeSolarizedLight:
		return &SolarizedLightTheme{}
	case ThemeWarm:
		return &WarmTheme{}
	case ThemeDark:
		fallthrough
	default:
		return NewAppTheme(theme.VariantDark)
	}
}

type AppTheme struct {
	// Mode explicitly forces VariantDark or VariantLight
	Mode fyne.ThemeVariant
}

func NewAppTheme(mode fyne.ThemeVariant) *AppTheme {
	return &AppTheme{Mode: mode}
}

// Color centralizes every single color look-up across the entire app
func (t *AppTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	activeVariant := t.Mode
	if activeVariant != theme.VariantLight && activeVariant != theme.VariantDark {
		activeVariant = variant
	}

	if activeVariant == theme.VariantLight {
		switch name {
		case theme.ColorNameBackground:
			return color.NRGBA{R: 245, G: 245, B: 250, A: 255}
		case theme.ColorNameForeground, theme.ColorNamePrimary:
			return color.NRGBA{R: 20, G: 20, B: 30, A: 255}
		case ColorNameSubtleText:
			return color.NRGBA{R: 100, G: 100, B: 115, A: 255}
		case ColorNameHoverCircle:
			return color.NRGBA{R: 0, G: 0, B: 0, A: 25}
		default:
			return theme.DefaultTheme().Color(name, theme.VariantLight)
		}
	}

	// Default: Dark Mode
	switch name {
	case theme.ColorNameBackground:
		return color.NRGBA{R: 24, G: 24, B: 36, A: 255}
	case theme.ColorNameForeground, theme.ColorNamePrimary:
		return color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	case ColorNameSubtleText:
		return color.NRGBA{R: 180, G: 180, B: 180, A: 255}
	case ColorNameHoverCircle:
		return color.NRGBA{R: 255, G: 255, B: 255, A: 35}
	default:
		return theme.DefaultTheme().Color(name, theme.VariantDark)
	}
}

func (t *AppTheme) Font(style fyne.TextStyle) fyne.Resource {
	if utils.InterFont != nil && len(utils.InterFont.StaticContent) > 0 {
		return utils.InterFont
	}
	return theme.DefaultTheme().Font(style)
}

func (t *AppTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t *AppTheme) Size(name fyne.ThemeSizeName) float32 {
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
