package controls

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func Msgbox(title, msg string, parent fyne.Window) {
	content := container.NewVBox(
		widget.NewLabel(msg),
	)
	custom := dialog.NewCustom(title, "dismiss", content, parent)
	custom.Resize(fyne.NewSize(240, 140))
	custom.SetDismissText(" 确定 ")
	custom.Show()
}
