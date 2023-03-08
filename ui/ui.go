package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewUI() fyne.CanvasObject {
	tabs := container.NewAppTabs(
		newHomeUI(),
		newSettingsUI(),
	)
	return tabs
}

// 更新UI的数据
func UpdateUIData() {
	linkTableUI.Update()
	linkPannelUI.Update()

	linkNewUI.Update()
	linkEditUI.Update()
	linkImportUI.Update()
	linkDeleteUI.Update()

	sysGeneralUI.Update()
	sysNetUI.Update()
	sysAboutUI.Update()
	sysBakUI.Update()
}

func AfterUIShowExec() {
	if GConfig.AutoLink {
		toolbarUI.start()
	}
}
