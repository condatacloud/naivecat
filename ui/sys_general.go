package ui

import (
	"fmt"
	"naivecat/tools"
	"naivecat/ui/controls"
	"naivecat/ui/recipe"
	"strconv"

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
	scaleEntry   *widget.Entry
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
	u.scaleEntry = widget.NewEntry()

	form1 := container.New(
		layout.NewFormLayout(),
		widget.NewLabel("主题"),
		u.themeSelect,
		widget.NewLabel("缩放"),
		u.scaleEntry,
	)
	vbox2 := container.NewVBox(
		u.enableLogCbx,
	)
	column := container.NewGridWithColumns(3,
		form1,
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
	scale, err := strconv.ParseFloat(u.scaleEntry.Text, 64)
	if err != nil {
		controls.Msgbox("错误", fmt.Sprintf("%s 不是正确的数字", u.scaleEntry.Text), Wnd)
		return
	}

	info := "成功保存配置"
	if u.theme != GConfig.Theme {
		info += "\n主题修改需要重新启动生效"
	}

	if !tools.IsEqual(scale, GConfig.Scale) {
		info += "\n缩放修改需要重新启动生效"
	}

	GConfig.Theme = u.theme
	GConfig.EnableLog = u.enableLog
	GConfig.Scale = scale
	GConfig.Update()
	controls.Msgbox("成功", info, Wnd)
}

func (u *SysGeneralUI) onCancelClicked() {
	u.Update()
}
