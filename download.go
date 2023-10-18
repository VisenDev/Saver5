package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

//this function creates the UI for the download menu
func DownloadMenu(model *Model, w *fyne.Window) fyne.CanvasObject {
	title := widget.NewRichTextFromMarkdown("# Download GCode From Machine")

	// input to enter a filepath
	input := widget.NewEntry()
	input.SetPlaceHolder("")

	return container.New(layout.NewVBoxLayout(), title, input)
}
