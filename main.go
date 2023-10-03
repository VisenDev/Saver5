package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"go.bug.st/serial"
)

type SerialConfig struct {
	Port     string
	Settings serial.Mode
}

func main() {
	a := app.New()
	w := a.NewWindow("Saver4")
	w.Resize(fyne.Size{800, 400})

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
		container.NewTabItem("Download", DownloadMenu(&configuration)),
	)

	tabs.SetTabLocation(container.TabLocationLeading)
	w.SetContent(tabs)
	w.ShowAndRun()
}

// TODO implement the Download Menu
func DownloadMenu(config *SerialConfig) fyne.CanvasObject {
	return widget.NewLabel("Download Menu")
}
