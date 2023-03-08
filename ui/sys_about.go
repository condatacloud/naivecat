package ui

import (
	"naivecat/model"
	"naivecat/service"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type SysAboutUI struct {
}

var sysAboutUI = &SysAboutUI{}

func (u *SysAboutUI) Update() {
}

func (u *SysAboutUI) NewUI() *fyne.Container {
	naiveVersion := service.NaiveService.GetVersion()
	vbox := container.NewVBox(
		container.NewHBox(widget.NewLabel("Git Commit    :"), widget.NewLabel(model.GitHash)),
		container.NewHBox(widget.NewLabel("Git Branch    :"), widget.NewLabel(model.GitBranch)),
		container.NewHBox(widget.NewLabel("Go Version    :"), widget.NewLabel(model.GoVersion)),
		container.NewHBox(widget.NewLabel("Build Time    :"), widget.NewLabel(model.BuildTime)),
		container.NewHBox(widget.NewLabel("Naive Version :"), widget.NewLabel(naiveVersion)),
	)
	return vbox
}
