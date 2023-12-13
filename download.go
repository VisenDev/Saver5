package main

import (
	"github.com/gen2brain/dlgs"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"fmt"
)

//this function creates the UI for the download menu
func DownloadMenu(model *Model, w *fyne.Window) fyne.CanvasObject {

	// Shows a preview of the file contents
	preview := widget.NewTextGrid()
	model.Listen(func() {
		fmt.Println("successfully called callback")
		preview.SetText(string(model.DownloadFileBuffer[0:model.DownloadBufLen]))
	})
	
	title := widget.NewRichTextFromMarkdown("# Recieved Text")

	clear_button := widget.NewButton("Clear Buffer", func(){
		model.DownloadBufLen = 0
		preview.SetText(string(model.DownloadFileBuffer[0:model.DownloadBufLen]))
	})


	//input := widget.NewEntry()
	//?/input.SetPlaceHolder("Enter filepath or select file...")
	
	save_button := widget.NewButton("Save To File", func(){
		filepath, _, err := dlgs.File("Select a file", "", false)

		if err != nil {
			DisplayError(w, "Failed to select file")
		} else {
			model.SetDownloadFilepath(filepath)
			model.WriteDownloadFile()
		}
	})

	return container.New(layout.NewVBoxLayout(), title, save_button, layout.NewSpacer(), clear_button, preview)
}
