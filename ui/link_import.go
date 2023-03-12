package ui

import (
	"naivecat/model"
	"naivecat/tools"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type LinkImportUI struct {
}

var linkImportUI = &LinkImportUI{}

func (u *LinkImportUI) Update() {
}

func (u *LinkImportUI) NewUI() {
	linkInput := widget.NewMultiLineEntry()
	linkInput.Wrapping = fyne.TextWrapBreak
	content := container.NewMax(linkInput)

	onClick := func(b bool) {
		if !b {
			return
		}

		u.importLink(tools.Strings.TrimN(linkInput.Text))
	}

	custom := dialog.NewCustomConfirm("导入链接", "确定", "取消", content, onClick, Wnd)
	custom.Resize(fyne.NewSize(400, 320))
	custom.Show()
}

func (u *LinkImportUI) importLink(txt string) {
	if txt == "" {
		return
	}
	link := &model.Link{}
	link.FromString(txt)
	GConfig.Links = append(GConfig.Links, link)
	GConfig.Update()
	linkTableUI.Update()
}
