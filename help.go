package main

import (
	"fyne.io/fyne/v2"
	//"fyne.io/fyne/v2/layout"
	//"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	//"go.bug.st/serial"
	//"log"
)

// this function returns the UI for the help menu
func HelpMenu() fyne.CanvasObject {
	return widget.NewLabel("This is the help menu!")
}
