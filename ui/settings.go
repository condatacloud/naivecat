package ui

import (
	"fyne.io/fyne/v2/container"
)

func newSettingsUI() *container.TabItem {
	tabs := container.NewAppTabs(
		container.NewTabItem("常规", sysGeneralUI.NewUI()),
		container.NewTabItem("网络", sysNetUI.NewUI()),
		container.NewTabItem("备份", sysBakUI.NewUI()),
		container.NewTabItem("关于", sysAboutUI.NewUI()),
	)
	tabs.SetTabLocation(container.TabLocationLeading)
	return container.NewTabItem("设置", tabs)
}
