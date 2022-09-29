package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/aakosarev/md5/internal"
	"strconv"
)

type keywordRestrictions struct {
	labelMaxLen         *widget.Label
	labelMinLen         *widget.Label
	labelMustContain    *widget.Label
	labelSpecialInclude *widget.Label
	maxLen              *widget.Entry
	minLen              *widget.Entry
	cancelBttn          *widget.Button
	applyBttn           *widget.Button
	checksRestriction   *widget.CheckGroup

	app    fyne.App
	window fyne.Window
}

func newKeywordRestrictions(app fyne.App, window fyne.Window) *keywordRestrictions {
	return &keywordRestrictions{
		app:    app,
		window: window,
	}
}

func (kr *keywordRestrictions) buildUI() *fyne.Container {

	kr.labelMaxLen = widget.NewLabel("Max length:")
	kr.labelMaxLen.Resize(fyne.NewSize(100, 10))
	kr.labelMaxLen.Move(fyne.NewPos(50, 30))

	kr.labelMinLen = widget.NewLabel("Min length:")
	kr.labelMinLen.Resize(fyne.NewSize(100, 10))
	kr.labelMinLen.Move(fyne.NewPos(50, 75))

	kr.labelMustContain = widget.NewLabel("The KeyWord \nmust contain:")
	kr.labelMustContain.Resize(fyne.NewSize(100, 30))
	kr.labelMustContain.Move(fyne.NewPos(50, 130))

	kr.labelSpecialInclude = widget.NewLabel("Special symbols include:\n\n\t\t\t ! ? # $ & % @ ^ * whitespace [ ]\n\t\t\t ( ) { } < > = ~ | _ ' \" \\ / : ; + - , . `")
	kr.labelSpecialInclude.Resize(fyne.NewSize(350, 30))
	kr.labelSpecialInclude.Move(fyne.NewPos(50, 290))

	kr.maxLen = widget.NewEntry()
	kr.maxLen.Resize(fyne.NewSize(200, 30))
	kr.maxLen.Move(fyne.NewPos(190, 35))
	maxLenKWStr := strconv.Itoa(internal.MaxLenKW)
	kr.maxLen.SetText(maxLenKWStr)

	kr.minLen = widget.NewEntry()
	kr.minLen.Resize(fyne.NewSize(200, 30))
	kr.minLen.Move(fyne.NewPos(190, 80))
	minLenKWStr := strconv.Itoa(internal.MinLenKW)
	kr.minLen.SetText(minLenKWStr)

	kr.cancelBttn = widget.NewButton("Cancel", func() {
		kr.window.Close()
	})
	kr.cancelBttn.Resize(fyne.NewSize(80, 30))
	kr.cancelBttn.Move(fyne.NewPos(245, 445))

	kr.applyBttn = widget.NewButton("Apply", func() {
		kr.apply()
	})
	kr.applyBttn.Resize(fyne.NewSize(80, 30))
	kr.applyBttn.Move(fyne.NewPos(120, 445))

	kr.checksRestriction = widget.NewCheckGroup([]string{"Digits", "Upper case", "Lower case", "Special symbols"}, func([]string) {})
	kr.checksRestriction.SetSelected(kr.getSelected())
	kr.checksRestriction.Move(fyne.NewPos(184, 130))

	return container.NewWithoutLayout(
		kr.labelMaxLen,
		kr.labelMinLen,
		kr.labelMustContain,
		kr.labelSpecialInclude,
		kr.maxLen,
		kr.minLen,
		kr.cancelBttn,
		kr.applyBttn,
		kr.checksRestriction,
	)
}

func (kr *keywordRestrictions) getSelected() []string {
	var selected []string

	if internal.DigitKW {
		selected = append(selected, "Digits")
	}
	if internal.UpperKW {
		selected = append(selected, "Upper case")
	}
	if internal.LowerKW {
		selected = append(selected, "Lower case")
	}
	if internal.SpecialKW {
		selected = append(selected, "Special symbols")
	}
	return selected
}

func (kr *keywordRestrictions) apply() {
	var err1, err2 error
	selected := kr.checksRestriction.Selected

	var minLenInt int
	minLenInt, err2 = strconv.Atoi(kr.minLen.Text)
	if err2 != nil {
		dialog.ShowInformation("Error", "Incorrect minimum keyword length.\n Enter the correct value!", kr.window)
		kr.minLen.SetText("")
		return
	}

	if minLenInt < len(selected) {
		minLemStr := strconv.Itoa(len(selected))
		kr.minLen.SetText(minLemStr)
	}

	internal.MaxLenKW, err1 = strconv.Atoi(kr.maxLen.Text)
	internal.MinLenKW, err2 = strconv.Atoi(kr.minLen.Text)
	if err1 != nil {
		dialog.ShowInformation("Error", "Incorrect maximum keyword length.\n Enter the correct value!", kr.window)
		kr.maxLen.SetText("")
	} else if err2 != nil {
		dialog.ShowInformation("Error", "Incorrect minimum keyword length.\n Enter the correct value!", kr.window)
		kr.minLen.SetText("")
	} else if internal.MinLenKW > internal.MaxLenKW {
		dialog.ShowInformation("Error", "The minimum length is less than the maximum.\n Enter the correct values!", kr.window)
		kr.maxLen.SetText("")
		kr.minLen.SetText("")
	} else {
		newSelected := kr.checksRestriction.Selected
		for _, v := range newSelected {
			switch v {
			case "Digits":
				internal.DigitKW = true
			case "Upper case":
				internal.UpperKW = true
			case "Lower case":
				internal.LowerKW = true
			case "Special symbols":
				internal.SpecialKW = true
			}
		}
		if !internal.Contains(newSelected, "Digits") {
			internal.DigitKW = false
		}
		if !internal.Contains(newSelected, "Upper case") {
			internal.UpperKW = false
		}
		if !internal.Contains(newSelected, "Lower case") {
			internal.LowerKW = false
		}
		if !internal.Contains(newSelected, "Special symbols") {
			internal.SpecialKW = false
		}
		kr.window.Close()
	}
}
