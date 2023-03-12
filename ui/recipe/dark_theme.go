package recipe

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type DarkTheme struct{}

func (t *DarkTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {

	switch name {
	case theme.ColorNamePrimary:
		return color.NRGBA{R: 0x29, G: 0x6f, B: 0xf6, A: 0xff}
	case theme.ColorNameFocus:
		return color.NRGBA{R: 0x00, G: 0xdb, B: 0x2a, A: 0x2a}
	case theme.ColorNameSelection:
		return color.NRGBA{R: 0, G: 163, B: 154, A: 255}
	}

	return t.paletColorNamed(name)
}

func (t *DarkTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t *DarkTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func (*DarkTheme) Font(s fyne.TextStyle) fyne.Resource {
	return fontResource
}

func (*DarkTheme) paletColorNamed(name fyne.ThemeColorName) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.NRGBA{R: 0x14, G: 0x14, B: 0x15, A: 0xff}
	case theme.ColorNameButton:
		return color.NRGBA{R: 0x28, G: 0x29, B: 0x2e, A: 0xff}
	case theme.ColorNameDisabled:
		return color.NRGBA{R: 0x39, G: 0x39, B: 0x3a, A: 0xff}
	case theme.ColorNameDisabledButton:
		return color.NRGBA{R: 0x28, G: 0x29, B: 0x2e, A: 0xff}
	case theme.ColorNameError:
		return errorColor
	case theme.ColorNameForeground:
		return color.NRGBA{R: 0xf3, G: 0xf3, B: 0xf3, A: 0xff}
	case theme.ColorNameHover:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x0f}
	case theme.ColorNameInputBackground:
		return color.NRGBA{R: 0x20, G: 0x20, B: 0x23, A: 0xff}
	case theme.ColorNameInputBorder:
		return color.NRGBA{R: 0x39, G: 0x39, B: 0x3a, A: 0xff}
	case theme.ColorNameMenuBackground:
		return color.NRGBA{R: 0x28, G: 0x29, B: 0x2e, A: 0xff}
	case theme.ColorNameOverlayBackground:
		return color.NRGBA{R: 0x18, G: 0x1d, B: 0x25, A: 0xff}
	case theme.ColorNamePlaceHolder:
		return color.NRGBA{R: 0xb2, G: 0xb2, B: 0xb2, A: 0xff}
	case theme.ColorNamePressed:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x66}
	case theme.ColorNameScrollBar:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x99}
	case theme.ColorNameSeparator:
		return color.NRGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xff}
	case theme.ColorNameShadow:
		return color.NRGBA{A: 0x66}
	case theme.ColorNameSuccess:
		return successColor
	case theme.ColorNameWarning:
		return warningColor
	}

	return color.Transparent
}
