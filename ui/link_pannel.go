package ui

import (
	"fmt"
	"naivecat/ui/controls"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type LinkPannelUI struct {
	name     *widget.Label
	host     *widget.Label
	port     *widget.Label
	protocol *widget.Label
	padding  *widget.Label
	qrcode   *canvas.Image
}

var linkPannelUI = &LinkPannelUI{}

func (u *LinkPannelUI) Update() {
	if len(GConfig.Links) == 0 {
		if u.qrcode != nil {
			u.name.SetText("")
			u.host.SetText("")
			u.port.SetText("")
			u.protocol.SetText("")
			u.padding.SetText("")
		}
		binding.NewString()
		return
	}
	if u.qrcode != nil {
		link := GConfig.Links[GConfig.DefaultLink]
		u.name.SetText(link.Name)
		u.host.SetText(link.Host)
		u.port.SetText(link.Port)
		u.protocol.SetText(link.Protocol)
		u.padding.SetText(fmt.Sprintf("%v", link.Padding))
		u.qrcode.Image = link.ToQCode()
		u.qrcode.Refresh()
	}
}

func (u *LinkPannelUI) NewUI() *fyne.Container {
	u.name = widget.NewLabel("                        ")
	u.host = widget.NewLabel("")
	u.port = widget.NewLabel("")
	u.protocol = widget.NewLabel("")
	u.padding = widget.NewLabel("")
	u.qrcode = &canvas.Image{}
	u.qrcode.FillMode = canvas.ImageFillOriginal
	u.qrcode.Resize(fyne.NewSize(168, 168))

	container.NewPadded()

	return container.NewHBox(
		container.NewVBox(
			container.NewHBox(widget.NewLabel("名称："), u.name),
			container.NewHBox(widget.NewLabel("地址："), u.host),
			container.NewHBox(widget.NewLabel("端口："), u.port),
			container.NewHBox(widget.NewLabel("协议："), u.protocol),
			container.NewHBox(widget.NewLabel("对齐："), u.padding),
		),
		layout.NewSpacer(),
		u.qrcode,
	)
}

func (u *LinkPannelUI) qrcode2Clipboard() {
	link := GConfig.Links[GConfig.DefaultLink]
	Wnd.Clipboard().SetContent(link.ToText())
	controls.Msgbox("信息", "成功复制链接到粘贴板", Wnd)
}
