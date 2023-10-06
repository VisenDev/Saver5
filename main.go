package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"go.bug.st/serial"
)

//this struct stores the configuration details that will be passed to the upload/download backend
type SerialConfig struct {
	Port     string
	Settings serial.Mode
}

func main() {
	a := app.New()

	//create a new window and resize it
	w := a.NewWindow("Saver4")
	w.Resize(fyne.Size{800, 400})

	//initializing the configuration happens here
	configuration := SerialConfig{
		Port: "",
		Settings: serial.Mode{
			BaudRate: 57600,
			Parity:   serial.EvenParity,
			DataBits: 7,
			StopBits: serial.OneStopBit,
		},
	}

	// create tabs for each of our windows
	tabs := container.NewAppTabs(
		container.NewTabItem("Serial Port Config", SerialSelectionMenu(&configuration)),
		container.NewTabItem("Upload", UploadMenu(&configuration, &w)),
		container.NewTabItem("Download", DownloadMenu(&configuration, &w)),
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	//set the content of our window to show the tabs
	w.SetContent(tabs)

	//render the window
	w.ShowAndRun()
}
