package ui

import (
	"fyne.io/fyne/v2/container"
)

func newHomeUI() *container.TabItem {
	linkTable := linkTableUI.NewUI()
	footerTools := toolbarUI.NewFooterUI()
	topTools := toolbarUI.NewTopUI()
	// 上下左右中间
	linkTableArea := container.NewBorder(nil, footerTools, nil, nil, linkTable)
	linkPannel := linkPannelUI.NewUI()
	console := consoleUI.NewUI()

	split := container.NewVSplit(
		linkPannel,
		console,
	)
	split.SetOffset(0.25)
	lay := container.NewBorder(topTools, nil, linkTableArea, nil, split)
	return container.NewTabItem("主页", lay)
}
