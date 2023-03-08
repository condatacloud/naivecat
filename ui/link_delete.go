package ui

import (
	"naivecat/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type LinkDeleteUI struct {
}

var linkDeleteUI = &LinkDeleteUI{}

func (u *LinkDeleteUI) Update() {
}

func (u *LinkDeleteUI) NewUI() {

	content := container.NewCenter(
		widget.NewLabel("确定删除该链接么?"),
	)

	onClick := func(b bool) {
		if !b {
			return
		}

		GConfig.Links = append(GConfig.Links[:GConfig.DefaultLink], GConfig.Links[GConfig.DefaultLink+1:]...)
		GConfig.DefaultLink = 0

		// 一个都没有生成默认的
		if len(GConfig.Links) == 0 {
			defaultLink := &model.Link{}
			defaultLink.NewDefaultLink()
			GConfig.Links = append(GConfig.Links, defaultLink)
		}

		GConfig.Update()
		linkTableUI.Update()
		linkPannelUI.Update()
	}
	custom := dialog.NewCustomConfirm("注意", "确定", "取消", content, onClick, Wnd)
	custom.Resize(fyne.NewSize(280, 180))
	custom.Show()
}
