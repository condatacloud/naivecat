package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("ProgressBar Widget")

	top := container.NewHBox(widget.NewLabel("sadda"), widget.NewButton("sdsds", func() {

	}))

	bottom := widget.NewMultiLineEntry()

	myWindow.SetContent(
		container.NewBorder(
			container.NewHBox(widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("Welcome")),
			nil,
			widget.NewLabel("Footer"),
			nil,

			container.NewBorder(
				top,
				nil,
				nil,
				nil,
				container.NewMax(bottom),
			),
		),
	)
	myWindow.ShowAndRun()
}
