package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type TransparentTheme struct{}

var _ fyne.Theme = (*TransparentTheme)(nil)

func (t TransparentTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameOverlayBackground {
		return color.NRGBA{R: 14, G: 14, B: 18, A: 255}
	}
	if name == theme.ColorNameBackground {
		return color.NRGBA{R: 12, G: 12, B: 16, A: 255}
	}
	if name == theme.ColorNameSeparator {
		return color.Transparent
	}
	if name == theme.ColorNameShadow {
		return color.Transparent
	}
	if name == theme.ColorNameInputBorder {
		return color.Transparent
	}
	if name == theme.ColorNameInputBackground {
		return color.NRGBA{R: 18, G: 18, B: 22, A: 255}
	}
	if name == theme.ColorNameMenuBackground {
		return color.NRGBA{R: 16, G: 16, B: 20, A: 255}
	}
	if name == theme.ColorNameScrollBarBackground {
		return color.Transparent
	}

	return theme.DarkTheme().Color(name, variant)
}

func (t TransparentTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DarkTheme().Font(style)
}

func (t TransparentTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DarkTheme().Icon(name)
}

func (t TransparentTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DarkTheme().Size(name)
}
