package resource

import (
	_ "embed"
)

//go:embed fonts/WenQuanDengKuanWeiMiHei-1.ttf
var ZhFontBytes []byte

//go:embed icon/naivecat.ico
var AppIcoBytes []byte

//go:embed icon/naivecat.png
var AppIconPngBytes []byte

//go:embed icon/internet.png
var IconInternet []byte

//go:embed icon/start.png
var IconStart []byte

//go:embed icon/stop.png
var IconStop []byte

//go:embed icon/share.png
var IconShare []byte
