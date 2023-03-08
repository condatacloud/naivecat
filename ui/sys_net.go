package ui

import (
	"naivecat/ui/controls"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type SysNetUI struct {
	hostEntry      *widget.Entry
	socksPortEntry *widget.Entry
	autoLinkCbx    *widget.Check
	httpEnableCbx  *widget.Check
	httpPortEntry  *widget.Entry
}

var sysNetUI = &SysNetUI{}

func (u *SysNetUI) Update() {
	u.hostEntry.SetText(GConfig.Host)
	u.socksPortEntry.SetText(GConfig.Socks.Port)
	u.autoLinkCbx.SetChecked(GConfig.AutoLink)
	u.httpEnableCbx.SetChecked(GConfig.Http.Enable)
	u.httpPortEntry.SetText(GConfig.Http.Port)
	if !GConfig.Http.Enable {
		u.httpPortEntry.Disable()
	}
}

func (u *SysNetUI) NewUI() *fyne.Container {
	u.hostEntry = widget.NewEntry()
	u.socksPortEntry = widget.NewEntry()
	u.autoLinkCbx = widget.NewCheck("自动链接", func(bool) {})
	u.httpEnableCbx = widget.NewCheck("启动http", u.onHttpEnableChanged)
	u.httpPortEntry = widget.NewEntry()
	u.httpPortEntry.SetPlaceHolder("http端口")

	vbox1 := container.NewVBox(
		container.New(
			layout.NewFormLayout(),
			widget.NewLabel("监听地址"),
			u.hostEntry,
			widget.NewLabel("socks端口"),
			u.socksPortEntry,
		),
	)
	vbox2 := container.NewVBox(
		u.autoLinkCbx,
		container.New(layout.NewFormLayout(), u.httpEnableCbx, u.httpPortEntry),
	)
	column := container.NewGridWithColumns(3,
		vbox1,
		vbox2,
	)

	okBtn := widget.NewButton("确定", u.onOkClicked)
	cancelBtn := widget.NewButton("取消", u.onCancelClicked)

	main := container.NewVBox(
		column,
		layout.NewSpacer(),
		container.NewHBox(layout.NewSpacer(), okBtn, cancelBtn),
	)

	u.Update()
	return main
}

func (u *SysNetUI) onHttpEnableChanged(b bool) {
	if !b {
		u.httpPortEntry.Disable()
	} else {
		u.httpPortEntry.Enable()
	}
}

func (u *SysNetUI) onOkClicked() {
	GConfig.Host = u.hostEntry.Text
	GConfig.Socks.Port = u.socksPortEntry.Text
	GConfig.AutoLink = u.autoLinkCbx.Checked
	GConfig.Http.Enable = u.httpEnableCbx.Checked
	GConfig.Http.Port = u.httpPortEntry.Text
	GConfig.Update()
	controls.Msgbox("成功", "成功保存配置", Wnd)
}

func (u *SysNetUI) onCancelClicked() {
	u.Update()
}
