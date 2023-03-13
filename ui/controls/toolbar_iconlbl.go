package controls

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ToolbarIconLbl struct {
	icon *widget.Icon
	lbl  *canvas.Text
}

func NewToolbarIconLbl(icon fyne.Resource, label string) *ToolbarIconLbl {
	i := widget.NewIcon(icon)
	i.MinSize()
	l := canvas.NewText(label, color.NRGBA{R: 0x43, G: 0xf4, B: 0x36, A: 0xff})
	l.MinSize()
	return &ToolbarIconLbl{icon: i, lbl: l}
}

func (t *ToolbarIconLbl) ToolbarObject() fyne.CanvasObject {
	return container.NewHBox(t.icon, container.NewCenter(t.lbl))
}

func (t *ToolbarIconLbl) SetText(txt string) {
	t.lbl.Text = txt
	t.lbl.Refresh()
}
