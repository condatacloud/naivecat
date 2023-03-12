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

	content := container.NewBorder(
		linkPannel,
		nil,
		nil,
		nil,
		container.NewMax(console),
	)

	lay := container.NewBorder(topTools, nil, linkTableArea, nil, content)
	return container.NewTabItem("主页", lay)
}
