package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/aakosarev/md5/internal"
)

type inputKeyword struct {
	labelKeyWord *widget.Label
	keyWord      *widget.Entry
	inputBttn    *widget.Button

	app    fyne.App
	window fyne.Window
}

var okk = false

func newInputKeyword(app fyne.App, window fyne.Window) *inputKeyword {
	return &inputKeyword{
		app:    app,
		window: window,
	}
}

func (ik *inputKeyword) buildUI() *fyne.Container {

	ik.labelKeyWord = widget.NewLabel("KeyWord:")
	ik.labelKeyWord.Resize(fyne.NewSize(100, 10))
	ik.labelKeyWord.Move(fyne.NewPos(30, 75))

	ik.keyWord = widget.NewPasswordEntry()
	ik.keyWord.SetPlaceHolder("Enter the KeyWord...")
	ik.keyWord.Resize(fyne.NewSize(200, 40))
	ik.keyWord.Move(fyne.NewPos(170, 75))

	ik.inputBttn = widget.NewButton("Enter", func() {
		internal.KWForCompare = ik.keyWord.Text
		okk = true
		ik.window.Close()
	})
	ik.inputBttn.Resize(fyne.NewSize(100, 40))
	ik.inputBttn.Move(fyne.NewPos(400, 75))

	return container.NewWithoutLayout(
		ik.labelKeyWord,
		ik.keyWord,
		ik.inputBttn,
	)
}
