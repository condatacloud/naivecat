package ui

import (
	"naivecat/ui/controls"
	"naivecat/ui/recipe"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type SysGeneralUI struct {
	theme        string
	enableLog    bool
	themeSelect  *widget.Select
	enableLogCbx *widget.Check
}

var sysGeneralUI = &SysGeneralUI{}

func (u *SysGeneralUI) Update() {
	u.theme = GConfig.Theme
	u.enableLog = GConfig.EnableLog
	u.themeSelect.SetSelected(u.theme)
	u.enableLogCbx.SetChecked(u.enableLog)
}

func (u *SysGeneralUI) NewUI() *fyne.Container {
	u.themeSelect = widget.NewSelect(recipe.THEMES, u.onThemeChanged)
	u.enableLogCbx = widget.NewCheck("开启日志", u.onEnableLogChanged)

	vbox1 := container.NewVBox(
		container.New(layout.NewFormLayout(), widget.NewLabel("主题"), u.themeSelect),
	)
	vbox2 := container.NewVBox(
		u.enableLogCbx,
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
	return main
}

func (u *SysGeneralUI) onThemeChanged(theme string) {
	u.theme = theme
}

func (u *SysGeneralUI) onEnableLogChanged(b bool) {
	u.enableLog = b
}

func (u *SysGeneralUI) onOkClicked() {
	info := "成功保存配置"
	if u.theme != GConfig.Theme {
		info += "\n主题需要重新启动生效"
	}
	GConfig.Theme = u.theme
	GConfig.EnableLog = u.enableLog
	GConfig.Update()
	controls.Msgbox("成功", info, Wnd)
}

func (u *SysGeneralUI) onCancelClicked() {
	u.Update()
}
