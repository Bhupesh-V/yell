package themes

import (
	"image/color"
	"sort"
	"yell/internal/ui/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

const (
	ColorNameSubtleText  fyne.ThemeColorName = "yell:color:subtleText"
	ColorNameHoverCircle fyne.ThemeColorName = "yell:color:hoverCircle"
)

type ThemeType string

// Centralized theme palette registry map
var themeRegistry = map[ThemeType]Palette{
	"dark": {
		Background:  color.NRGBA{R: 24, G: 24, B: 36, A: 255},
		Foreground:  color.NRGBA{R: 255, G: 255, B: 255, A: 255},
		SubtleText:  color.NRGBA{R: 180, G: 180, B: 180, A: 255},
		HoverCircle: color.NRGBA{R: 255, G: 255, B: 255, A: 35},
		Variant:     theme.VariantDark,
	},
	"light": {
		Background:  color.NRGBA{R: 245, G: 245, B: 250, A: 255},
		Foreground:  color.NRGBA{R: 20, G: 20, B: 30, A: 255},
		SubtleText:  color.NRGBA{R: 100, G: 100, B: 115, A: 255},
		HoverCircle: color.NRGBA{R: 0, G: 0, B: 0, A: 25},
		Variant:     theme.VariantLight,
	},
	"warm": {
		Background:  color.NRGBA{R: 222, G: 198, B: 149, A: 255},
		Foreground:  color.NRGBA{R: 42, G: 31, B: 26, A: 255},
		SubtleText:  color.NRGBA{R: 105, G: 82, B: 71, A: 255},
		HoverCircle: color.NRGBA{R: 42, G: 31, B: 26, A: 35},
		Variant:     theme.VariantLight,
	},
	"nord": {
		Background:  color.NRGBA{R: 46, G: 52, B: 64, A: 255},
		Foreground:  color.NRGBA{R: 236, G: 239, B: 244, A: 255},
		SubtleText:  color.NRGBA{R: 143, G: 188, B: 187, A: 255},
		HoverCircle: color.NRGBA{R: 216, G: 222, B: 233, A: 35},
		Variant:     theme.VariantDark,
	},
	"forest": {
		Background:  color.NRGBA{R: 28, G: 38, B: 30, A: 255},
		Foreground:  color.NRGBA{R: 218, G: 228, B: 212, A: 255},
		SubtleText:  color.NRGBA{R: 138, G: 160, B: 132, A: 255},
		HoverCircle: color.NRGBA{R: 163, G: 190, B: 140, A: 30},
		Variant:     theme.VariantDark,
	},
	"cyberpunk": {
		Background:  color.NRGBA{R: 9, G: 2, B: 33, A: 255},
		Foreground:  color.NRGBA{R: 0, G: 255, B: 204, A: 255},
		SubtleText:  color.NRGBA{R: 255, G: 105, B: 180, A: 255},
		HoverCircle: color.NRGBA{R: 0, G: 255, B: 204, A: 40},
		Variant:     theme.VariantDark,
	},
	"solarized-dark": {
		Background:  color.NRGBA{R: 0, G: 43, B: 54, A: 255},
		Foreground:  color.NRGBA{R: 131, G: 148, B: 150, A: 255},
		SubtleText:  color.NRGBA{R: 101, G: 123, B: 131, A: 255},
		HoverCircle: color.NRGBA{R: 7, G: 54, B: 66, A: 180},
		Variant:     theme.VariantDark,
	},
	"solarized-light": {
		Background:  color.NRGBA{R: 253, G: 246, B: 227, A: 255},
		Foreground:  color.NRGBA{R: 88, G: 110, B: 117, A: 255},
		SubtleText:  color.NRGBA{R: 147, G: 161, B: 161, A: 255},
		HoverCircle: color.NRGBA{R: 238, G: 232, B: 213, A: 180},
		Variant:     theme.VariantLight,
	},
	"dracula": {
		Background:  color.NRGBA{R: 40, G: 42, B: 54, A: 255},
		Foreground:  color.NRGBA{R: 248, G: 248, B: 242, A: 255},
		SubtleText:  color.NRGBA{R: 189, G: 147, B: 249, A: 255},
		HoverCircle: color.NRGBA{R: 255, G: 121, B: 198, A: 40},
		Variant:     theme.VariantDark,
	},
	"catppuccin": {
		Background:  color.NRGBA{R: 30, G: 30, B: 46, A: 255},
		Foreground:  color.NRGBA{R: 205, G: 214, B: 244, A: 255},
		SubtleText:  color.NRGBA{R: 147, G: 153, B: 178, A: 255},
		HoverCircle: color.NRGBA{R: 137, G: 180, B: 250, A: 40},
		Variant:     theme.VariantDark,
	},
	"tokyo-night": {
		Background:  color.NRGBA{R: 26, G: 27, B: 38, A: 255},
		Foreground:  color.NRGBA{R: 192, G: 202, B: 245, A: 255},
		SubtleText:  color.NRGBA{R: 122, G: 162, B: 247, A: 255},
		HoverCircle: color.NRGBA{R: 238, G: 117, B: 254, A: 40},
		Variant:     theme.VariantDark,
	},
	"monokai": {
		Background:  color.NRGBA{R: 45, G: 42, B: 46, A: 255},
		Foreground:  color.NRGBA{R: 255, G: 216, B: 102, A: 255},
		SubtleText:  color.NRGBA{R: 120, G: 220, B: 232, A: 255},
		HoverCircle: color.NRGBA{R: 255, G: 97, B: 136, A: 40},
		Variant:     theme.VariantDark,
	},
	// Gruvbox Dark (Warm Retro)
	"gruvbox": {
		Background:  color.NRGBA{R: 40, G: 40, B: 40, A: 255},    // Dark medium gray
		Foreground:  color.NRGBA{R: 251, G: 241, B: 199, A: 255}, // Cream / Pale yellow
		SubtleText:  color.NRGBA{R: 254, G: 128, B: 25, A: 255},  // Warm orange accent
		HoverCircle: color.NRGBA{R: 251, G: 241, B: 199, A: 30},
		Variant:     theme.VariantDark,
	},

	// Matrix / Hacker Green (Monochrome Terminal)
	"matrix": {
		Background:  color.NRGBA{R: 10, G: 15, B: 10, A: 255}, // Ultra-dark green tint
		Foreground:  color.NRGBA{R: 0, G: 255, B: 65, A: 255}, // Bright terminal green
		SubtleText:  color.NRGBA{R: 0, G: 143, B: 17, A: 255}, // Dimmed phosphor green
		HoverCircle: color.NRGBA{R: 0, G: 255, B: 65, A: 35},
		Variant:     theme.VariantDark,
	},

	// Synthwave '84 (High-contrast Neon Dark)
	"synthwave": {
		Background:  color.NRGBA{R: 38, G: 20, B: 51, A: 255},    // Deep purple night
		Foreground:  color.NRGBA{R: 255, G: 222, B: 89, A: 255},  // Neon yellow title
		SubtleText:  color.NRGBA{R: 255, G: 126, B: 219, A: 255}, // Hot pink text
		HoverCircle: color.NRGBA{R: 54, G: 241, B: 205, A: 40},   // Cyan hover
		Variant:     theme.VariantDark,
	},

	// OLED Black (True Pitch Black for AMOLED Displays)
	"oled-black": {
		Background:  color.NRGBA{R: 0, G: 0, B: 0, A: 255},       // Pitch black
		Foreground:  color.NRGBA{R: 255, G: 255, B: 255, A: 255}, // Sharp white
		SubtleText:  color.NRGBA{R: 160, G: 160, B: 160, A: 255}, // Neutral gray
		HoverCircle: color.NRGBA{R: 255, G: 255, B: 255, A: 40},
		Variant:     theme.VariantDark,
	},

	// Coffee / Espresso (Rich Dark Browns & Caramel)
	"espresso": {
		Background:  color.NRGBA{R: 30, G: 20, B: 18, A: 255},    // Dark roasted coffee bean
		Foreground:  color.NRGBA{R: 240, G: 220, B: 200, A: 255}, // Cream / Foam text
		SubtleText:  color.NRGBA{R: 210, G: 150, B: 90, A: 255},  // Warm caramel accent
		HoverCircle: color.NRGBA{R: 240, G: 220, B: 200, A: 30},
		Variant:     theme.VariantDark,
	},

	// Sunset / Vaporwave (Coral, Soft Gold, and Deep Plum)
	"sunset": {
		Background:  color.NRGBA{R: 45, G: 20, B: 44, A: 255},  // Deep plum sky
		Foreground:  color.NRGBA{R: 255, G: 183, B: 3, A: 255}, // Sun gold title text
		SubtleText:  color.NRGBA{R: 251, G: 133, B: 0, A: 255}, // Warm coral secondary text
		HoverCircle: color.NRGBA{R: 255, G: 183, B: 3, A: 35},
		Variant:     theme.VariantDark,
	},

	// Cyberpunk Gold / 2077 (High-Contrast Yellow & Slate)
	"cyberpunk-2077": {
		Background:  color.NRGBA{R: 20, G: 20, B: 20, A: 255},  // Near-black dark slate
		Foreground:  color.NRGBA{R: 243, G: 230, B: 0, A: 255}, // Signature Cyberpunk yellow
		SubtleText:  color.NRGBA{R: 0, G: 220, B: 255, A: 255}, // Bright electric cyan text
		HoverCircle: color.NRGBA{R: 243, G: 230, B: 0, A: 45},
		Variant:     theme.VariantDark,
	},

	// Everforest (Soft Comforting Warm Dark Green)
	"everforest": {
		Background:  color.NRGBA{R: 43, G: 51, B: 57, A: 255},    // Muted dark gray-green
		Foreground:  color.NRGBA{R: 211, G: 198, B: 170, A: 255}, // Soft warm beige text
		SubtleText:  color.NRGBA{R: 167, G: 192, B: 128, A: 255}, // Soft olive green text
		HoverCircle: color.NRGBA{R: 211, G: 198, B: 170, A: 30},
		Variant:     theme.VariantDark,
	},

	// Classic Paper / Sepia (Soft Warm Light Theme)
	"sepia": {
		Background:  color.NRGBA{R: 244, G: 236, B: 216, A: 255}, // Aged paper background
		Foreground:  color.NRGBA{R: 67, G: 52, B: 34, A: 255},    // Dark ink brown text
		SubtleText:  color.NRGBA{R: 140, G: 110, B: 80, A: 255},  // Medium brown secondary text
		HoverCircle: color.NRGBA{R: 67, G: 52, B: 34, A: 25},
		Variant:     theme.VariantLight,
	},
}

// ParseTheme converts string input directly into ThemeType with a fallback lookup
func ParseTheme(input string) ThemeType {
	t := ThemeType(input)
	if _, exists := themeRegistry[t]; exists {
		return t
	}
	return "dark"
}

// GetTheme resolves the theme directly from the map
func GetTheme(t ThemeType) fyne.Theme {
	if palette, exists := themeRegistry[t]; exists {
		return NewBaseTheme(palette)
	}
	return NewBaseTheme(themeRegistry["dark"])
}

// Themes returns all available themes
func Themes() []string {
	names := make([]string, 0, len(themeRegistry))
	for k := range themeRegistry {
		names = append(names, string(k))
	}
	sort.Strings(names)
	return names
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
