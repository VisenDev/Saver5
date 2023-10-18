package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"go.bug.st/serial"
	"log"
)

func SerialSelectionMenu(m *Model) fyne.CanvasObject {

	//generate the title widget
	title := widget.NewRichTextFromMarkdown("# Serial Configuration")
	port_selector_title := widget.NewLabel("Please Select a Port")

	port_selector := CreateSerialSelection(m)
   
   refresh_button := widget.NewButton("Refresh Serial Ports", func() {
      port_selector = CreateSerialSelection(m)
	})

	return container.New(layout.NewVBoxLayout(), title, port_selector_title, port_selector, refresh_button)
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
