package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"sync"
)

var (
	once     sync.Once
	instance *App
)

type App struct {
	app    fyne.App
	window fyne.Window
}

func GetApp() *App {
	once.Do(func() {
		a := app.NewWithID("123")
		w := a.NewWindow("MD5")
		instance = &App{
			app:    a,
			window: w,
		}
	})
	return instance
}

func (a *App) Run() {
	a.window.Resize(fyne.NewSize(705, 560))
	a.app.Settings().SetTheme(theme.DarkTheme())
	icon, _ := fyne.LoadResourceFromPath("./internal/icons/icon.png")
	a.window.SetIcon(icon)
	a.window.SetFixedSize(true)
	a.window.CenterOnScreen()
	a.window.SetMaster()
	a.window.SetContent(newMain(a.app, a.window).buildUI())
	a.window.Show()
	a.app.Run()
}
