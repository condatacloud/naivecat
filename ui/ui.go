package ui

import (
	"naivecat/ui/recipe"

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

func InitTheme() {
	var theme fyne.Theme
	// 设置主题、字体
	if GConfig.Theme == recipe.THEME_DARK {
		theme = &recipe.DarkTheme{}
	} else if GConfig.Theme == recipe.THEME_LIGHT {
		theme = &recipe.LightTheme{}
	} else {
		theme = &recipe.DarkTheme{}
	}
	App.Settings().SetTheme(theme)
}

// 更新UI的数据
func InitUIData() {
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
