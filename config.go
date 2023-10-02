package main 

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"go.bug.st/serial"
	"log"
)

func SerialSelectionMenu(config *SerialConfig) fyne.CanvasObject {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
      return widget.NewLabel("Fatal Error: No serial ports found!")
	}

   //generate the title widget
	port_selector_title := widget.NewLabel("Please Select a Serial Port")
   
   //generate the list of buttons to select por	//}
   port_selector := widget.NewSelect(ports, func(value string) {
      config.Port = value
	})

    return container.New(layout.NewVBoxLayout(), port_selector_title, port_selector)
}
