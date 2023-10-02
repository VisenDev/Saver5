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

   selected_port := -1
   tabs := container.NewAppTabs(
		container.NewTabItem("Serial Port", SerialSelectionMenu(&selected_port)),
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)

	//tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))
	tabs.SetTabLocation(container.TabLocationLeading)

	//current_menu := SerialSelectionMenu(w)
	w.SetContent(tabs)
	w.ShowAndRun()
}

func SerialSelectionMenu(selected_port *int) fyne.CanvasObject {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}

   //generate the title widget
	title := widget.NewLabel("Please Select a Serial Port")
   
   //generate the widget to display which port is currently selected
   display_text := "none"
   if *selected_port >= 0 && *selected_port < len(ports) {
      display_text = ports[*selected_port]
   }
   selected_display := widget.NewLabel("Currently Selected: " + display_text)
   
   //combine title and display
   info := container.NewGridWithRows(2)
   info.Add(title)
   info.Add(selected_display)

   //generate the list of buttons to select port
	port_button_list := container.NewGridWithRows(len(ports) + 1)
	for i := 0; i < len(ports); i++ {
      current_port_id := i
		button := widget.NewButton(ports[i], func() {
         *selected_port = current_port_id
		})
		port_button_list.Add(button)
	}

	content := container.NewGridWithColumns(2)
   content.Add(info)
   content.Add(port_button_list)
   
	return content 
}
