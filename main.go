package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"go.bug.st/serial"
	"log"
)

func main() {
	a := app.New()
	w := a.NewWindow("Saver4")
	w.Resize(fyne.Size{800, 400})

	content := showSerialPorts(w)
	w.SetContent(content)
	w.ShowAndRun()
}

func showSerialPorts(w fyne.Window) fyne.CanvasObject {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}

	list := container.NewGridWithRows(len(ports) + 1)
	title := widget.NewLabel("Please Select a Serial Port")

	for i := 0; i < len(ports); i++ {
      current_port := ports[i]
		button := widget.NewButton(ports[i], func() {
         message := widget.NewLabel("Serial Port Selected: " + current_port)
         w.SetContent(message)
		})
		list.Add(button)
	}

	content := container.NewGridWithColumns(2)
   content.Add(title)
   content.Add(list)

	return content 
}
