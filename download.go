package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func DownloadMenu(config *SerialConfig, w *fyne.Window) fyne.CanvasObject {
	title := widget.NewRichTextFromMarkdown("# Download GCode From Machine")

	// input to enter a filepath
	input := widget.NewEntry()
	input.SetPlaceHolder("")

	return container.New(layout.NewVBoxLayout(), title, input)
}
