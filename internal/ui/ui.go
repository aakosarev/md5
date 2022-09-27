package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"sync"
	"time"
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

	driver := a.app.Driver().(desktop.Driver)
	splash := driver.CreateSplashWindow()
	splash.SetContent(
		container.NewVBox(
			layout.NewSpacer(),
			widget.NewLabelWithStyle("Loading...", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
			layout.NewSpacer(),
			layout.NewSpacer(),
			widget.NewProgressBarInfinite(),
			layout.NewSpacer(),
		),
	)
	go func() {
		time.Sleep(time.Second * 3)
		a.window.Resize(fyne.NewSize(705, 560))
		a.window.SetFixedSize(true)
		a.window.CenterOnScreen()
		a.window.SetMaster()
		a.window.SetContent(a.setupUI())
		a.window.Show()

		splash.Close()
	}()

	splash.ShowAndRun()
}

func (a *App) setupUI() fyne.CanvasObject {

	return &container.AppTabs{Items: []*container.TabItem{
		uiMain(a.app, a.window).tabItem(),
	}}
}
