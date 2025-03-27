package ui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func Window() {
	skunkDatingApp := app.New()
	w := skunkDatingApp.NewWindow("Skunk Dating App")
	w.SetContent(widget.NewLabel("Welcome to the Skunk Dating App!"))
	w.Show()
	skunkDatingApp.Run()
}
