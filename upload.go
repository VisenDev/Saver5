package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func UploadMenu(config *SerialConfig, w *fyne.Window) fyne.CanvasObject {
	title := widget.NewRichTextFromMarkdown("# Upload GCode")
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter filepath or select file...")

	// path to the chosen file
	var filepath string = ""

	filepicker_button := widget.NewButton("Select a file to upload", func() {
		filepicker := dialog.NewFileOpen(func(url fyne.URIReadCloser, err error) {
			if err != nil {
				log.Fatal(err)
			}
			if url != nil {
				filepath = url.URI().Path()
				input.SetText(filepath)
			}
		}, *w)
		filepicker.Show()
	})

	// file contents holds the preview
	preview := widget.NewRichTextFromMarkdown("")

	preview_button := widget.NewButton("Preview Selected File", func() {
		if filepath != "" {
			bytes, err := os.ReadFile(filepath)
			if err != nil {
				log.Fatal(err)
			}
			preview.ParseMarkdown("```\n" + string(bytes) + "\n```")
		}
	})

	return container.New(layout.NewVBoxLayout(), title, input, filepicker_button, preview_button, preview)
}
