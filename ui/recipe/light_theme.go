package recipe

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type LightTheme struct{}

func (t *LightTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNamePrimary:
		return color.NRGBA{R: 0x29, G: 0x6f, B: 0xf6, A: 0xff}
	case theme.ColorNameFocus:
		return color.NRGBA{R: 0x00, G: 0xdb, B: 0x2a, A: 0x2a}
	case theme.ColorNameSelection:
		return color.NRGBA{R: 0x00, G: 0x36, B: 0x40, A: 0x40}
	}

	return t.paletColorNamed(name)
}

func (t *LightTheme) Icon(name fyne.ThemeIconName) fyne.Resource {

	return theme.DefaultTheme().Icon(name)
}

func (t *LightTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func (*LightTheme) Font(s fyne.TextStyle) fyne.Resource {
	return fontResource
}

func (*LightTheme) paletColorNamed(name fyne.ThemeColorName) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	case theme.ColorNameButton:
		return color.NRGBA{R: 0xf5, G: 0xf5, B: 0xf5, A: 0xff}
	case theme.ColorNameDisabled:
		return color.NRGBA{R: 0xe3, G: 0xe3, B: 0xe3, A: 0xff}
	case theme.ColorNameDisabledButton:
		return color.NRGBA{R: 0xf5, G: 0xf5, B: 0xf5, A: 0xff}
	case theme.ColorNameError:
		return errorColor
	case theme.ColorNameForeground:
		return color.NRGBA{R: 0x56, G: 0x56, B: 0x56, A: 0xff}
	case theme.ColorNameHover:
		return color.NRGBA{A: 0x0f}
	case theme.ColorNameInputBackground:
		return color.NRGBA{R: 0xf3, G: 0xf3, B: 0xf3, A: 0xff}
	case theme.ColorNameInputBorder:
		return color.NRGBA{R: 0xe3, G: 0xe3, B: 0xe3, A: 0xff}
	case theme.ColorNameMenuBackground:
		return color.NRGBA{R: 0xf5, G: 0xf5, B: 0xf5, A: 0xff}
	case theme.ColorNameOverlayBackground:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	case theme.ColorNamePlaceHolder:
		return color.NRGBA{R: 0x88, G: 0x88, B: 0x88, A: 0xff}
	case theme.ColorNamePressed:
		return color.NRGBA{A: 0x19}
	case theme.ColorNameScrollBar:
		return color.NRGBA{A: 0x99}
	case theme.ColorNameSeparator:
		return color.NRGBA{R: 0xf5, G: 0xf5, B: 0xf5, A: 0xff}
	case theme.ColorNameShadow:
		return color.NRGBA{A: 0x33}
	case theme.ColorNameSuccess:
		return successColor
	case theme.ColorNameWarning:
		return warningColor
	}

	return color.Transparent
}
