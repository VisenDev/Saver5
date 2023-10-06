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

//this function returns the UI for the upload menu
func UploadMenu(config *SerialConfig, w *fyne.Window) fyne.CanvasObject {
	title := widget.NewRichTextFromMarkdown("# Upload GCode")

	// input to enter a filepath
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter filepath or select file...")

	// button to pick a file
	filepicker_button := widget.NewButton("Select a file to upload", func() {
		filepicker := dialog.NewFileOpen(func(url fyne.URIReadCloser, err error) {
			if err != nil {
				log.Fatal(err)
			}
			if url != nil {
				filepath := url.URI().Path()
				input.SetText(filepath)
			}
		}, *w)
		//show filepicker if button is pressed
		filepicker.Show()
	})

	// Shows a preview of the file contents
	preview := widget.NewRichTextFromMarkdown("")

	//contain the preview in a container with a scroll bar
	preview_wrapper := container.NewScroll(preview)
	preview_wrapper.SetMinSize(fyne.Size{100, 300})

	//logic for the preview selected file button
	preview_button := widget.NewButton("Preview Selected File", func() {
		filepath := input.Text
		if filepath != "" {
			bytes, err := os.ReadFile(filepath)
			if err != nil {
				popup := container.New(layout.NewCenterLayout(), widget.NewRichTextFromMarkdown("# failed to open file"))
				widget.ShowPopUp(popup, (*w).Canvas())
			}
			preview.ParseMarkdown("```\n" + string(bytes) + "\n```")
		} else {
			// show error sayning that input text is input
			popup := container.New(layout.NewCenterLayout(), widget.NewRichTextFromMarkdown("# no file chosen"))
			widget.ShowPopUp(popup, (*w).Canvas())
		}
	})

	//logic for Upload
	upload_button := widget.NewButton("Upload!", func() {})

	button_wrapper := container.New(layout.NewGridLayout(3), filepicker_button, preview_button, upload_button)

	//combine all of the widgets and return
	return container.New(layout.NewVBoxLayout(), title, input, button_wrapper, preview_wrapper)
}
