package main

import (
   "log"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
   "fyne.io/fyne/v2"
   "fyne.io/fyne/v2/layout"
   "fyne.io/fyne/v2/container"
	"go.bug.st/serial"
)

func main() {
	a := app.New()
	w := a.NewWindow("Saver4")

	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}

//   
//	for _, port := range ports {
//      w.SetContent(widget.NewLabel(port))
//	}

	list := widget.NewList(
		func() int {
			return len(ports)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(ports[i])
		})

   title := widget.NewLabel("All Available Serial Ports")
   content := container.New(layout.NewHBoxLayout(), title, list, layout.NewSpacer())  
   
   w.SetContent(content)
	w.ShowAndRun()
}
