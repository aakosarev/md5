package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/aakosarev/md5/internal"
	"github.com/aakosarev/md5/internal/md5"
	"io"
)

type main struct {
	labelText            *widget.Label
	input                *widget.Entry
	labelHash            *widget.Label
	hash                 *widget.Entry
	setKeyWordBttn       *widget.Button
	checkSaveHashToFile  *widget.Check
	checkGetHashFromFile *widget.Check
	checkGetTextFromFile *widget.Check
	checkUseKeyWord      *widget.Check
	getHashBttn          *widget.Button
	compareBttn          *widget.Button
	menu                 *fyne.MainMenu

	app    fyne.App
	window fyne.Window
}

func newMain(app fyne.App, window fyne.Window) *main {
	return &main{
		app:    app,
		window: window,
	}
}

func (m *main) buildUI() *fyne.Container {
	m.labelText = widget.NewLabel("Text")
	m.labelText.Resize(fyne.NewSize(600, 10))
	m.labelText.Move(fyne.NewPos(45, 10))

	m.input = widget.NewMultiLineEntry()
	m.input.PlaceHolder = "Enter the text..."
	m.input.Wrapping = fyne.TextWrapBreak
	m.input.Resize(fyne.NewSize(600, 200))
	m.input.Move(fyne.NewPos(50, 40))

	m.labelHash = widget.NewLabel("Hash")
	m.labelHash.Resize(fyne.NewSize(600, 10))
	m.labelHash.Move(fyne.NewPos(45, 370))

	m.hash = widget.NewMultiLineEntry()
	m.hash.Wrapping = fyne.TextWrapBreak
	m.hash.Resize(fyne.NewSize(600, 80))
	m.hash.Move(fyne.NewPos(50, 400))

	m.setKeyWordBttn = widget.NewButton("Set KeyWord", func() {
		wait := true
		setKeyWordWindow := m.app.NewWindow("Set KeyWord")
		setKeyWordWindow.Resize(fyne.NewSize(490, 245))
		setKeyWordWindow.SetFixedSize(true)
		setKeyWordWindow.CenterOnScreen()
		sk := newSetKeyword(m.app, setKeyWordWindow)
		setKeyWordWindow.SetContent(sk.buildUI())
		setKeyWordWindow.Show()
		setKeyWordWindow.SetOnClosed(func(){
			wait = false
		})
		for wait {
		}
	})
	m.setKeyWordBttn.Resize(fyne.NewSize(130, 40))
	m.setKeyWordBttn.Move(fyne.NewPos(50, 315))

	m.checkSaveHashToFile = widget.NewCheck("Save hash to File", func(c bool) {})
	m.checkSaveHashToFile.Resize(fyne.NewSize(150, 40))
	m.checkSaveHashToFile.Move(fyne.NewPos(255, 290))

	m.checkGetHashFromFile = widget.NewCheck("Get Hash from File", func(c bool) {})
	m.checkGetHashFromFile.Resize(fyne.NewSize(150, 60))
	m.checkGetHashFromFile.Move(fyne.NewPos(485, 245))

	m.checkGetTextFromFile = widget.NewCheck("Get text from File", func(c bool) {})
	m.checkGetTextFromFile.Resize(fyne.NewSize(150, 40))
	m.checkGetTextFromFile.Move(fyne.NewPos(255, 245))

	m.checkUseKeyWord = widget.NewCheck("Use KeyWord", func(c bool) {})
	m.checkUseKeyWord.Resize(fyne.NewSize(150, 40))
	m.checkUseKeyWord.Move(fyne.NewPos(255, 335))

	m.getHashBttn = widget.NewButton("Get Hash", func() {
		var data []byte

		if m.checkGetTextFromFile.Checked {
			fopenDialog := dialog.NewFileOpen(
				func(file fyne.URIReadCloser, err error) {
					defer func() {
						if file != nil {
							file.Close()
						}
					}()
					if err != nil {
						return
					}
					if file == nil {
						return
					}
					data, _ = io.ReadAll(file)
					m.input.SetText(string(data))
					if m.checkUseKeyWord.Checked {
						data = append(data, []byte(internal.KW)...)
					}
					hash := md5.CalcMD5(data)
					m.hash.SetText(hash)
					if m.checkSaveHashToFile.Checked {
						m.saveDataToFile(hash)
					}
				}, m.window)
			fopenDialog.Resize(fyne.NewSize(500, 500))
			fopenDialog.Show()
		} else {
			data = []byte(m.input.Text)
			if m.checkUseKeyWord.Checked {
				data = append(data, []byte(internal.KW)...)
			}
			hash := md5.CalcMD5(data)
			m.hash.SetText(hash)
			if m.checkSaveHashToFile.Checked {
				m.saveDataToFile(hash)
			}
		}

	})
	m.getHashBttn.Resize(fyne.NewSize(130, 40))
	m.getHashBttn.Move(fyne.NewPos(50, 260))

	m.compareBttn = widget.NewButton("Compare", func() {
		var dataHash []byte
		if m.checkGetHashFromFile.Checked {
			fopenDialog := dialog.NewFileOpen(
				func(file fyne.URIReadCloser, err error) {
					defer func() {
						if file != nil {
							file.Close()
						}
					}()
					if err != nil {
						return
					}
					if file == nil {
						return
					}
					dataHash, _ = io.ReadAll(file)

					var data []byte
					if m.checkGetTextFromFile.Checked {
						m.compareIfTextFromFile(&data, dataHash)
					} else {
						data = []byte(m.input.Text)
						if m.checkUseKeyWord.Checked {
							m.inputKeyWordForCompare()
							data = append(data, []byte(internal.KWForCompare)...)
						}
						hash := md5.CalcMD5(data)
						m.compareHashs(hash, string(dataHash))
					}
				}, m.window)
			fopenDialog.Resize(fyne.NewSize(500, 500))
			fopenDialog.Show()
		} else {
			dataHash = []byte(m.hash.Text)
			var data []byte
			if m.checkGetTextFromFile.Checked {
				m.compareIfTextFromFile(&data, dataHash)
			} else {
				data = []byte(m.input.Text)
				if m.checkUseKeyWord.Checked {
					m.inputKeyWordForCompare()
					data = append(data, []byte(internal.KWForCompare)...)
				}
				hash := md5.CalcMD5(data)
				m.compareHashs(hash, string(dataHash))
			}
		}
	})
	m.compareBttn.Resize(fyne.NewSize(150, 60))
	m.compareBttn.Move(fyne.NewPos(500, 300))

	menuItemSaveTextToFile := fyne.NewMenuItem("Save text to File", func() {
		m.saveDataToFile(m.input.Text)
	})
	menuFile := fyne.NewMenu("File", menuItemSaveTextToFile)
	menuItemSetKeyWordRestrictions := fyne.NewMenuItem("Set a KeyWord restriction", func() {
		wait := true
		keyWordRestrictionsWindow := m.app.NewWindow("KeyWord restriction")
		keyWordRestrictionsWindow.Resize(fyne.NewSize(460, 500))
		keyWordRestrictionsWindow.SetFixedSize(true)
		keyWordRestrictionsWindow.CenterOnScreen()
		kr := newKeywordRestrictions(m.app, keyWordRestrictionsWindow)
		keyWordRestrictionsWindow.SetContent(kr.buildUI())
		keyWordRestrictionsWindow.Show()
		keyWordRestrictionsWindow.SetOnClosed(func(){
			wait = false
		})
		for wait {
		}
	})
	menuOptions := fyne.NewMenu("Options", menuItemSetKeyWordRestrictions)
	menuItemAbout := fyne.NewMenuItem("About", func() {
		dialog.ShowInformation(
			"About",
			"Косарев А.А.\n\nА-05-19\n\nПрограммная реализация функции хеширования MD5",
			m.window,
		)
	})
	menuHelp := fyne.NewMenu("Help", menuItemAbout)

	m.menu = fyne.NewMainMenu(menuFile, menuOptions, menuHelp)
	m.window.SetMainMenu(m.menu)

	return container.NewWithoutLayout(
		m.labelText,
		m.input,
		m.labelHash,
		m.hash,
		m.setKeyWordBttn,
		m.checkSaveHashToFile,
		m.checkGetHashFromFile,
		m.checkGetTextFromFile,
		m.checkUseKeyWord,
		m.getHashBttn,
		m.compareBttn,
	)
}

func (m *main) saveDataToFile(data string) {
	fsDialog := dialog.NewFileSave(
		func(file fyne.URIWriteCloser, err error) {
			if file == nil {
				return
			}
			if err != nil {
				return
			}
			dataInBytes := []byte(data)
			_, _ = file.Write(dataInBytes)
		},
		m.window,
	)
	fsDialog.Resize(fyne.NewSize(500, 500))
	fsDialog.Show()
}

func (m *main) compareHashs(newHash string, oldHash string) {
	if newHash == string(oldHash) {
		dialog.ShowInformation("Success", "Hashes matched!", m.window)
	} else {
		dialog.ShowInformation("Fail", "Hashes not matched!", m.window)
	}
}

func (m *main) compareIfTextFromFile(data *[]byte, dataHash []byte) {
	fopenDialog := dialog.NewFileOpen(
		func(file fyne.URIReadCloser, err error) {
			defer func() {
				if file != nil {
					file.Close()
				}
			}()
			if err != nil {
				return
			}
			if file == nil {
				return
			}
			*data, _ = io.ReadAll(file)
			if m.checkUseKeyWord.Checked {
				m.inputKeyWordForCompare()
				*data = append(*data, []byte(internal.KWForCompare)...)
			}
			hash := md5.CalcMD5(*data)
			m.compareHashs(hash, string(dataHash))
		}, m.window)
	fopenDialog.Resize(fyne.NewSize(500, 500))
	fopenDialog.Show()
}

func (m *main) inputKeyWordForCompare(){
	wait := true
	inputKeyWordWindow := m.app.NewWindow("Enter of the KeyWord")
	inputKeyWordWindow.Resize(fyne.NewSize(550, 210))
	inputKeyWordWindow.SetFixedSize(true)
	inputKeyWordWindow.CenterOnScreen()
	ck := newInputKeyword(m.app, inputKeyWordWindow)
	inputKeyWordWindow.SetContent(ck.buildUI())
	inputKeyWordWindow.Show()
	inputKeyWordWindow.SetOnClosed(func() {
		wait = false
	})
	for wait {
	}
}
