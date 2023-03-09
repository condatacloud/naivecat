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

// 图标链接
// https://fonts.google.com/icons?selected=Material+Symbols+Outlined:share:FILL@1;wght@300;GRAD@0;opsz@48&icon.platform=web

// 注意！！！！
// svg图标必须带有 fill=\"white\" fill="white" 和 viewBox="0 0 48 48"
var IconNetwork = []byte("<svg xmlns=\"http://www.w3.org/2000/svg\" height=\"48\" width=\"48\" viewBox=\"0 0 48 48\"><path fill=\"white\" d=\"M3.45 21.45 1.1 19.1q4.4-4.5 10.35-7.025Q17.4 9.55 24 9.55q1.05 0 2.45.075 1.4.075 2.65.225l-1.55 3.2-1.725-.1Q24.95 12.9 24 12.9q-5.9 0-11.175 2.275T3.45 21.45Zm6.75 6.6L7.85 25.7q3-2.9 7.125-4.525Q19.1 19.55 24.55 19.65L23.1 22.7q-3.65.05-7.15 1.575-3.5 1.525-5.75 3.775Zm12.35 11.2q-1.2-.5-1.7-1.725t.1-2.525L31.9 11.65q.1-.3.425-.45.325-.15.625-.05.35.1.55.4.2.3.1.6L26.8 37.4q-.35 1.35-1.675 1.825-1.325.475-2.575.025Zm15.25-11.2q-.75-.8-2-1.725t-2.25-1.375l.8-3.2q1.5.55 3.125 1.725T40.15 25.7Zm6.75-6.6q-1.8-1.8-4.075-3.35-2.275-1.55-4.375-2.55l.85-3.35q3.05 1.35 5.4 3.025Q44.7 16.9 46.9 19.1Z\"/></svg>")

var IconStart = []byte("<svg xmlns=\"http://www.w3.org/2000/svg\" height=\"48\" width=\"48\" viewBox=\"0 0 48 48\"><path fill=\"white\" d=\"M16 37.85v-28l22 14Z\"/></svg>")

var IconStop = []byte("<svg xmlns=\"http://www.w3.org/2000/svg\" height=\"48\" width=\"48\" viewBox=\"0 0 48 48\"><path fill=\"white\" d=\"M26.25 38V10H38v28ZM10 38V10h11.75v28Z\"/></svg>")

var IconShared = []byte("<svg xmlns=\"http://www.w3.org/2000/svg\" height=\"48\" width=\"48\" viewBox=\"0 0 48 48\"><path d=\"M36.25 43.5q-2.2 0-3.725-1.55T31 38.2q0-.35.075-.85t.225-.85l-15.6-9.1q-.7.85-1.75 1.375t-2.15.525q-2.2 0-3.75-1.55Q6.5 26.2 6.5 24t1.55-3.75Q9.6 18.7 11.8 18.7q1.1 0 2.125.475T15.7 20.5l15.6-9q-.15-.4-.225-.85Q31 10.2 31 9.8q0-2.2 1.525-3.75Q34.05 4.5 36.25 4.5T40 6.05q1.55 1.55 1.55 3.7 0 2.25-1.55 3.775t-3.75 1.525q-1.15 0-2.175-.4T32.35 13.4L16.8 22.2q.1.4.175.9.075.5.075.9t-.075.825q-.075.425-.175.825l15.55 8.9q.7-.7 1.675-1.15.975-.45 2.225-.45 2.2 0 3.75 1.525Q41.55 36 41.55 38.2T40 41.95q-1.55 1.55-3.75 1.55Z\"/></svg>")
