package main

import (
	"time"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

//this function creates the UI for the download menu
func DownloadMenu(model *Model, w *fyne.Window) fyne.CanvasObject {

	// Shows a preview of the file contents
	preview := widget.NewTextGrid()
	
	listen_button := widget.NewButton("Wait for input", func(){
		if model.Port == nil {
			DisplayError(w, "no open port")
			return
		}
		
		buff := make([]byte, 100000)
		sum := 0

		t, _ := time.ParseDuration("10s")
		_ = model.Port.SetReadTimeout(t)
		
		for {
			n, err := model.Port.Read(buff)
			if err != nil {
				DisplayError(w, "error reading from port")
				return
			}
			if n == 0 {
				break
			}
			sum += n
			fmt.Println(string(buff))
			model.DownloadFileBuffer = string(buff[0:sum])
			preview.SetText(model.DownloadFileBuffer)
		}
	})

	title := widget.NewRichTextFromMarkdown("# Recieved Text")

	return container.New(layout.NewVBoxLayout(), listen_button, title, preview)
}
