package recipe

import (
	"image/color"
	"naivecat/resource"

	"fyne.io/fyne/v2"
)

var fontResource *fyne.StaticResource

func init() {
	fontResource = fyne.NewStaticResource("WenQuanDengKuanWeiMiHei", resource.ZhFontBytes)
}

var (
	errorColor   = color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff}
	successColor = color.NRGBA{R: 0x43, G: 0xf4, B: 0x36, A: 0xff}
	warningColor = color.NRGBA{R: 0xff, G: 0x98, B: 0x00, A: 0xff}
)

const (
	THEME_LIGHT = "Light"
	THEME_DARK  = "Dark"
)

var (
	THEMES = []string{THEME_LIGHT, THEME_DARK}
)
