package ui

import (
	"naivecat/service"
	"naivecat/tools"
	"naivecat/ui/recipe"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type SetupExecFunc func(func(float32))

func NewSetupUI(title string, execFunc SetupExecFunc) {
	setup := app.New()
	setup.Settings().SetTheme(&recipe.DarkTheme{})

	progressBar := widget.NewProgressBar()

	vbox := container.NewVBox(
		widget.NewLabel(title),
		progressBar,
	)

	drv := setup.Driver()
	if drv, ok := drv.(desktop.Driver); ok {
		wnd := drv.CreateSplashWindow()

		update := func(v float32) {
			if tools.IsEqual(float64(v), 1.0) {
				time.Sleep(1 * time.Second)
				progressBar.SetValue(1)
				time.Sleep(1400 * time.Millisecond)
				wnd.Close()
			} else {
				progressBar.SetValue(float64(v))
			}
		}

		wnd.SetContent(vbox)
		wnd.Show()

		go execFunc(update)

		setup.Run()
	}
}

func NewUpgradeUI() {
	NewSetupUI(
		"      Naivecat升级      ",
		service.SetupService.Upgrade,
	)
}

func NewInstallUI() {
	NewSetupUI(
		"      Naivecat安装      ",
		service.SetupService.Install,
	)
}
