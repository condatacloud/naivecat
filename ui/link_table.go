package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type LinkTableUI struct {
	data  []string
	pings []string
	list  *widget.List
}

var linkTableUI = &LinkTableUI{}

func (u *LinkTableUI) Update() {
	u.data = GConfig.Links.ToNameList()
	u.pings = GConfig.Links.ToPingList()
	u.list.Refresh()
	u.list.Select(GConfig.DefaultLink)
}

func (u *LinkTableUI) NewUI() *widget.List {
	u.list = widget.NewList(
		func() int { return len(u.data) }, //最终显示个数
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewLabel("123456789012345678901"), layout.NewSpacer(), widget.NewLabel("1000ms"))
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*fyne.Container).Objects[0].(*widget.Label).SetText(u.data[i])
			o.(*fyne.Container).Objects[2].(*widget.Label).SetText(u.pings[i])
		},
	)
	u.list.OnSelected = func(id widget.ListItemID) {
		if len(GConfig.Links) == 0 {
			return
		}
		GConfig.DefaultLink = id
		linkPannelUI.Update()
	}

	return u.list
}
