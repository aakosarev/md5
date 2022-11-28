package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/aakosarev/md5/internal"
)

type setKeyword struct {
	labelKeyWord      *widget.Label
	labelConfirmation *widget.Label
	keyWord           *widget.Entry
	confirmation      *widget.Entry
	cancelBttn        *widget.Button
	setBttn           *widget.Button

	app    fyne.App
	window fyne.Window
}

var ok = false

func newSetKeyword(app fyne.App, window fyne.Window) *setKeyword {
	return &setKeyword{
		app:    app,
		window: window,
	}
}

func (sk *setKeyword) buildUI() *fyne.Container {

	sk.labelKeyWord = widget.NewLabel("KeyWord:")
	sk.labelKeyWord.Resize(fyne.NewSize(100, 10))
	sk.labelKeyWord.Move(fyne.NewPos(30, 37))

	sk.labelConfirmation = widget.NewLabel("Confirmation:")
	sk.labelConfirmation.Resize(fyne.NewSize(100, 10))
	sk.labelConfirmation.Move(fyne.NewPos(30, 95))

	sk.keyWord = widget.NewPasswordEntry()
	sk.keyWord.PlaceHolder = "Enter the KeyWord..."
	sk.keyWord.Resize(fyne.NewSize(265, 40))
	sk.keyWord.Move(fyne.NewPos(170, 35))

	sk.confirmation = widget.NewPasswordEntry()
	sk.confirmation.PlaceHolder = "Enter the KeyWord again..."
	sk.confirmation.Resize(fyne.NewSize(265, 40))
	sk.confirmation.Move(fyne.NewPos(170, 95))

	sk.cancelBttn = widget.NewButton("Cancel", func() {
		sk.window.Close()
	})
	sk.cancelBttn.Resize(fyne.NewSize(100, 40))
	sk.cancelBttn.Move(fyne.NewPos(260, 185))

	sk.setBttn = widget.NewButton("Set", func() {
		if sk.keyWord.Text != sk.confirmation.Text {
			dialog.ShowInformation("Error", "KeyWords don't match!\nRe-enter!", sk.window)
			sk.keyWord.SetText("")
			sk.confirmation.SetText("")
		} else if !internal.CheckKeyWord(sk.keyWord.Text) {
			infoDialog := dialog.NewInformation("Error", "The KeyWord does not meet the restrictions.\nSee the \"Options\" tab!", sk.window)
			sk.keyWord.SetText("")
			sk.confirmation.SetText("")
			infoDialog.SetOnClosed(func() {
				sk.window.Close()
			})
			infoDialog.Show()

		} else {
			internal.KW = sk.keyWord.Text
			ok = true
			sk.window.Close()
		}
	})
	sk.setBttn.Resize(fyne.NewSize(100, 40))
	sk.setBttn.Move(fyne.NewPos(120, 185))

	return container.NewWithoutLayout(
		sk.labelKeyWord,
		sk.labelConfirmation,
		sk.keyWord,
		sk.confirmation,
		sk.cancelBttn,
		sk.setBttn,
	)
}
