package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// HelpMenu returns the UI for the help menu.

func AboutMenu() fyne.CanvasObject {
	header := widget.NewLabel("KAGER INDUSTRIES - SAVER5")
	page_content := widget.NewLabel("Written by:\n-Robert Burnett\n\nCo-Designed by:\n-Connor Bumann\n-Elizabeth Salyards\n-Eduardo")

	return container.NewVBox(
		header,
		page_content,
	)
}
