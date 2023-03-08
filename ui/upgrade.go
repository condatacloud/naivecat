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

func NewUpgradeUI() {
	upgrade := app.New()
	upgrade.Settings().SetTheme(&recipe.DarkTheme{})

	progressBar := widget.NewProgressBar()

	vbox := container.NewVBox(
		widget.NewLabel("      Naivecat升级      "),
		progressBar,
	)

	drv := upgrade.Driver()
	if drv, ok := drv.(desktop.Driver); ok {
		wnd := drv.CreateSplashWindow()

		update := func(v float32) {
			if tools.IsEqual(float64(v), 1.0) {
				time.Sleep(1 * time.Second)
				progressBar.SetValue(1)
				time.Sleep(1 * time.Second)
				wnd.Close()
			} else {
				progressBar.SetValue(float64(v))
			}
		}

		wnd.SetContent(vbox)
		wnd.Show()

		go service.SetupService.Upgrade(update)

		upgrade.Run()
	}
}
