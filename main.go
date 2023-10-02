package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"go.bug.st/serial"
	"log"
)

func main() {
	a := app.New()
	w := a.NewWindow("Saver4")
	w.Resize(fyne.Size{800, 400})

   selected_port := -1
   
   //create tabs for each of our windows
   tabs := container.NewAppTabs(
		container.NewTabItem("Serial Port Config", SerialSelectionMenu(&selected_port)),
		container.NewTabItem("Upload", UploadMenu()),
		container.NewTabItem("Download", DownloadMenu()),
	)

	tabs.SetTabLocation(container.TabLocationLeading)
	w.SetContent(tabs)
	w.ShowAndRun()
}

//TODO implement the Upload Menu
func UploadMenu() fyne.CanvasObject {
    return widget.NewLabel("Upload Menu")  
}

//TODO implement the Download Menu
func DownloadMenu() fyne.CanvasObject {
    return widget.NewLabel("Download Menu")  
}


//TODO make a struct for containing serial port configurations
func SerialSelectionMenu(selected_port *int) fyne.CanvasObject {
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
		log.Println("Port Selected", value)
      //TODO have this update the selected port
	})

    return container.New(layout.NewVBoxLayout(), port_selector_title, port_selector)
}
