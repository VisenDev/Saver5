package main

import (
//	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"go.bug.st/serial"
	"log"
)

func SerialSelectionMenu(m *Model,  w *fyne.Window) fyne.CanvasObject {

	//generate the title widget
	//title := widget.NewRichTextFromMarkdown("# Serial Configuration")
	//port_selector_title := widget.NewLabel("Please Select a Port")
	port_selector := CreateSerialSelection(m)
	selector := container.New(layout.NewMaxLayout(), port_selector)

	progress_bar := widget.NewProgressBarInfinite()
	progress_bar.Stop()
	connection_label := widget.NewLabel("Not Connected");
	connection_status := container.New(layout.NewVBoxLayout(), connection_label, progress_bar);
	
	connect_button := widget.NewButton("Check Connection", func() {
		var err error
		m.Port , err = serial.Open(m.Config.Port, &m.Config.Settings)	
		//buf := make([]byte, 10000)
		//m.Port.Read(buf)
		//fmt.Println(string(buf))
		if err != nil {
			DisplayError(w, "Failed to open port")
			progress_bar.Stop()
			connection_label.SetText("Not Connected")
			return
		}
		progress_bar.Start()
		connection_label.SetText("Connected to " + m.Config.Port)
	})

	var content fyne.CanvasObject
	var refresh_button fyne.CanvasObject
   
   refresh_button = widget.NewButton("Refresh Serial Ports", func() {
		port_selector = nil
      port_selector = CreateSerialSelection(m)
		content = container.New(layout.NewVBoxLayout(), selector, refresh_button, connect_button)
	})

	content = container.New(layout.NewVBoxLayout(), selector, refresh_button, connect_button)
	return container.New(layout.NewVBoxLayout(), content, layout.NewSpacer(), connection_status)
}

// this function returns the UI for the configuration menu
func CreateSerialSelection(m *Model) fyne.CanvasObject {
   ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		return widget.NewLabel("Fatal Error: No serial ports found!")
	}
   //generate the list of buttons to select por	//
	port_selector := widget.NewSelect(ports, func(value string) {
		m.Config.Port = value
	})
   
   return port_selector
}
