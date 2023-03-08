package recipe

import (
	"image/color"
	"naivecat/resource"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var fontResource *fyne.StaticResource

func init() {
	fontResource = fyne.NewStaticResource("WenQuanDengKuanWeiMiHei", resource.ZhFontBytes)
}

type DarkTheme struct{}

func (t *DarkTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	// switch name {
	// case theme.ColorNameBackground:
	// case theme.ColorNameButton:
	// case theme.ColorNameDisabledButton:
	// case theme.ColorNameDisabled:
	// case theme.ColorNameError:
	// case theme.ColorNameFocus:
	// case theme.ColorNameForeground:
	// case theme.ColorNameHover:
	// case theme.ColorNameInputBackground:
	// case theme.ColorNameInputBorder:
	// case theme.ColorNameMenuBackground:
	// case theme.ColorNameOverlayBackground:
	// case theme.ColorNamePlaceHolder:
	// case theme.ColorNamePressed:
	// case theme.ColorNamePrimary:
	// case theme.ColorNameScrollBar:
	// case theme.ColorNameSelection:
	// case theme.ColorNameSeparator:
	// case theme.ColorNameShadow:
	// case theme.ColorNameSuccess:
	// case theme.ColorNameWarning:
	// }

	// theme.LightTheme()

	return theme.DefaultTheme().Color(name, variant)
}

func (t *DarkTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t *DarkTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func (*DarkTheme) Font(s fyne.TextStyle) fyne.Resource {
	// if s.Monospace {
	// 	return theme.DefaultTheme().Font(s)
	// }
	// // 此处可以根据不同字重返回不同的字体, 但是我用的都是同样的字体
	// if s.Bold {
	// 	if s.Italic {
	// 		return theme.DefaultTheme().Font(s)
	// 	}
	// 	// 返回自定义字体
	// 	return fontResource
	// }
	// if s.Italic {
	// 	return theme.DefaultTheme().Font(s)
	// }
	// 返回自定义字体
	return fontResource
}
