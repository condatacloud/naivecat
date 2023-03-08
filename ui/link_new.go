package ui

import (
	"naivecat/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type LinkNewUI struct {
	link model.Link
}

var linkNewUI = &LinkNewUI{}

func (u *LinkNewUI) Update() {
}

func (u *LinkNewUI) NewUI() {
	nameEntry := widget.NewEntry()
	hostEntry := widget.NewEntry()
	portEntry := widget.NewEntry()
	usernameEntry := widget.NewEntry()
	passwordEntry := widget.NewEntry()
	protocolSelect := widget.NewSelect(model.LinkProtocols, u.onProtocolChanged)
	protocolSelect.SetSelected("https")
	paddingCbx := widget.NewCheck("", u.onPaddingChanged)

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
		u.addLink()
	}

	custom := dialog.NewCustomConfirm("新建链接", "确定", "取消", form, onClick, Wnd)
	custom.Resize(fyne.NewSize(420, 320))
	custom.Show()
}

func (u *LinkNewUI) addLink() {
	newLink := &model.Link{}
	newLink.Copy(&u.link)
	GConfig.Links = append(GConfig.Links, newLink)
	GConfig.Update()
	linkTableUI.Update()
}

func (u *LinkNewUI) onProtocolChanged(p string) {
	u.link.Protocol = p
}

func (u *LinkNewUI) onPaddingChanged(b bool) {
	u.link.Padding = b
}
