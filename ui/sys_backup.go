package ui

import (
	"naivecat/model"
	"naivecat/service"
	"naivecat/tools"
	"naivecat/ui/controls"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/sqweek/dialog"
)

type SysBakUI struct {
}

var sysBakUI = &SysBakUI{}

func (u *SysBakUI) Update() {}

func (u *SysBakUI) NewUI() *fyne.Container {
	vbox := container.NewVBox(
		widget.NewButton("导入备份", u.onImportBackup),
		widget.NewButton("导出备份", u.onExportBackup),
	)

	return container.NewWithoutLayout(vbox)
}

func (u *SysBakUI) onImportBackup() {
	if service.NaiveService.IsRunning() {
		controls.Msgbox("提示", "请先关闭运行的服务", Wnd)
		return
	}

	filename, err := dialog.File().Filter("选择备份文件", "ncbak").Load()
	if err != nil {
		controls.Msgbox("错误", err.Error(), Wnd)
		return
	}

	filePath := filename

	conf := &model.Config{}

	if err := tools.Deserialize(conf, filePath); err != nil {
		controls.Msgbox("错误", err.Error(), Wnd)
		return
	}

	GConfig = conf
	GConfig.Update()
	InitUIData()
	controls.Msgbox("成功", "导入备份成功", Wnd)
}

func (u *SysBakUI) onExportBackup() {
	directory, err := dialog.Directory().Title("导出目录").Browse()
	if err != nil {
		controls.Msgbox("错误", err.Error(), Wnd)
		return
	}

	filePath := directory + "/naivecat.ncbak"
	if err := tools.Serialize(GConfig, filePath); err != nil {
		controls.Msgbox("错误", err.Error(), Wnd)
		return
	}
	controls.Msgbox("成功", "导出备份成功", Wnd)
}
