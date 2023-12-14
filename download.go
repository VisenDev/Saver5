package main

import (
	//"github.com/gen2brain/dlgs"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"fmt"
	"github.com/sqweek/dialog"
)

//this function creates the UI for the download menu
func DownloadMenu(model *Model, w *fyne.Window) fyne.CanvasObject {

	// Shows a preview of the file contents
	preview := widget.NewTextGrid()

	// contain the preview in a container with a scroll bar
	preview_scroll := container.NewScroll(preview)
	preview_scroll.SetMinSize(fyne.Size{100, 300})
	
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
		//filepath, _, err := dlgs.File("Select a file", "", false)
		file := dialog.File()
		filepath, err := file.Save()

		if err != nil {
			DisplayError(w, "Failed to select file")
		} else {
			model.SetDownloadFilepath(filepath)
			model.WriteDownloadFile()
		}
	})

	button_wrapper := container.New(layout.NewGridLayout(2), save_button, clear_button)
	return container.New(layout.NewVBoxLayout(), title, button_wrapper, preview_scroll)
}
