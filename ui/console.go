package ui

import (
	"bytes"
	"container/list"

	"fyne.io/fyne/v2/widget"
)

type ConsoleUI struct {
	txt       *list.List
	cache     *bytes.Buffer
	multiLine *widget.Entry
}

var consoleUI = &ConsoleUI{
	txt:   list.New(),
	cache: bytes.NewBufferString(""),
}

func (u *ConsoleUI) Update() {
	if u.multiLine == nil {
		return
	}

	u.cache.Reset()
	for i := u.txt.Front(); i != nil; i = i.Next() {
		u.cache.WriteString(i.Value.(string))
		u.cache.WriteString("\n")
	}

	u.multiLine.SetText(u.cache.String())
}

func (u *ConsoleUI) append(s string) {
	if u.txt.Len() > 300 {
		u.txt.Remove(u.txt.Front())
	}
	u.txt.PushBack(s)
	u.Update()
}

func (u *ConsoleUI) clear() {
	var next *list.Element
	for e := u.txt.Front(); e != nil; e = next {
		next = e.Next()
		u.txt.Remove(e)
	}
	u.Update()
}

func (u *ConsoleUI) NewUI() *widget.Entry {
	u.multiLine = widget.NewMultiLineEntry()
	// u.multiLine.Disable()
	return u.multiLine
}
