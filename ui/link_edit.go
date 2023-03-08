package ui

import (
	"naivecat/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type LinkEditUI struct {
	link *model.Link
}

var linkEditUI = &LinkEditUI{
	link: &model.Link{},
}

func (u *LinkEditUI) Update() {
}

func (u *LinkEditUI) NewUI() {
	if len(GConfig.Links) == 0 {
		return
	}

	link := GConfig.Links[GConfig.DefaultLink]
	u.link.Copy(link)

	nameEntry := widget.NewEntry()
	hostEntry := widget.NewEntry()
	portEntry := widget.NewEntry()
	usernameEntry := widget.NewEntry()
	passwordEntry := widget.NewEntry()
	protocolSelect := widget.NewSelect(model.LinkProtocols, u.onProtocolChanged)
	paddingCbx := widget.NewCheck("", u.onPaddingChanged)

	nameEntry.SetText(u.link.Name)
	hostEntry.SetText(u.link.Host)
	portEntry.SetText(u.link.Port)
	usernameEntry.SetText(u.link.Username)
	passwordEntry.SetText(u.link.Password)
	protocolSelect.SetSelected(u.link.Protocol)
	paddingCbx.SetChecked(u.link.Padding)

	form := container.New(
		layout.NewFormLayout(),
		widget.NewLabel("链接名称"),
		nameEntry,
		widget.NewLabel("主机地址"),
		hostEntry,
		widget.NewLabel("主机端口"),
		portEntry,
		widget.NewLabel("用户名称"),
		usernameEntry,
		widget.NewLabel("用户密码"),
		passwordEntry,
		widget.NewLabel("传输协议"),
		protocolSelect,
		widget.NewLabel("数据对齐"),
		paddingCbx,
	)

	onClick := func(b bool) {
		if !b {
			return
		}

		u.link.Name = nameEntry.Text
		u.link.Host = hostEntry.Text
		u.link.Port = portEntry.Text
		u.link.Username = usernameEntry.Text
		u.link.Password = passwordEntry.Text
		u.changeLink()
	}
	custom := dialog.NewCustomConfirm("编辑链接", "确定", "取消", form, onClick, Wnd)
	custom.Resize(fyne.NewSize(420, 320))
	custom.Show()
}

func (u *LinkEditUI) onProtocolChanged(p string) {
	u.link.Protocol = p
}

func (u *LinkEditUI) onPaddingChanged(b bool) {
	u.link.Padding = b
}

func (u *LinkEditUI) changeLink() {
	link := GConfig.Links[GConfig.DefaultLink]
	link.Copy(u.link)
	GConfig.Update()
	linkTableUI.Update()
	linkPannelUI.Update()
}
