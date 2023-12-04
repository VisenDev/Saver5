package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/gen2brain/dlgs"
)

// this function returns the UI for the upload menu
func UploadMenu(model *Model, w *fyne.Window) fyne.CanvasObject {
	title := widget.NewRichTextFromMarkdown("# Upload GCode")

	// input to enter a filepath
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter filepath or select file...")

	// button to pick a file
	filepicker_button := widget.NewButton("Select a file to upload", func() {
		filepath, _, err := dlgs.File("Select a file", "", false)

		if err != nil {
			DisplayError(w, "Failed to select file")
		} else {
			input.SetText(filepath)
		}
	})

	// Shows a preview of the file contents
	preview := widget.NewTextGrid()

	// contain the preview in a container with a scroll bar
	preview_wrapper := container.NewScroll(preview)
	preview_wrapper.SetMinSize(fyne.Size{100, 300})

	// logic for the preview selected file button
	preview_button := widget.NewButton("Preview Selected File", func() {

		model.SetUploadFilepath(input.Text)
		file_contents, err := model.ReadUploadFile()

		if err != nil {
			DisplayError(w, "Failed to read file")
		} else {
			// set contents to loaded file
			//preview.ParseMarkdown("```\n" + file_contents + "\n```")
			preview.SetText(file_contents)
		}
	})

	// logic for Upload
	upload_button := widget.NewButton("Upload!", func() {
		err := model.SetUploadFilepath(input.Text)
		if err != nil {
			DisplayError(w, "failed to parse filepath")
		}
		_, err = model.Upload(w)
		if err != nil {
			DisplayError(w, "failed to upload file")
		}
	})

	button_wrapper := container.New(layout.NewGridLayout(3), filepicker_button, preview_button, upload_button)

	// combine all of the widgets and return
	return container.New(layout.NewVBoxLayout(), title, input, button_wrapper, preview_wrapper)
}
