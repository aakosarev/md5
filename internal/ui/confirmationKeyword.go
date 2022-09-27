package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/aakosarev/md5/internal"
)

type confirmationKeyword struct {
	labelKeyWord *widget.Label
	keyWord      *widget.Entry
	confirmBttn  *widget.Button

	app    fyne.App
	window fyne.Window
}

func newConfirmationKeyword(app fyne.App, window fyne.Window) *confirmationKeyword {
	return &confirmationKeyword{
		app:    app,
		window: window,
	}
}

func (ck *confirmationKeyword) buildUI() *fyne.Container {

	ck.labelKeyWord = widget.NewLabel("KeyWord:")
	ck.labelKeyWord.Resize(fyne.NewSize(100, 10))
	ck.labelKeyWord.Move(fyne.NewPos(30, 75))

	ck.keyWord = widget.NewPasswordEntry()
	ck.keyWord.SetPlaceHolder("Enter the KeyWord...")
	ck.keyWord.Resize(fyne.NewSize(200, 40))
	ck.keyWord.Move(fyne.NewPos(170, 75))

	ck.confirmBttn = widget.NewButton("Confirm", func() {
		if ck.keyWord.Text == internal.KW {
			internal.Confirmed = true
			infoDialog := dialog.NewInformation("Success", "The KeyWord has been successfully confirmed", ck.window)
			infoDialog.SetOnClosed(func() {
				ck.window.Close()
			})
			infoDialog.Show()
		} else {
			internal.Confirmed = false
			infoDialog := dialog.NewInformation("Error", "Wrong KeyWord!", ck.window)
			infoDialog.SetOnClosed(func() {
				ck.window.Close()
			})
			infoDialog.Show()
		}
	})
	ck.confirmBttn.Resize(fyne.NewSize(100, 40))
	ck.confirmBttn.Move(fyne.NewPos(400, 75))

	return container.NewWithoutLayout(
		ck.labelKeyWord,
		ck.keyWord,
		ck.confirmBttn,
	)
}
