package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// this function returns the UI for the help menu
func HelpMenu() fyne.CanvasObject {
	title := widget.NewRichTextFromMarkdown("# Help")
	sub_title_1 :=  widget.NewRichTextFromMarkdown( "## My serial cable is not showing up on Windows" )
	content_1 := widget.NewRichTextFromMarkdown("- In order to use this application with an usb-to-rs232 adapter on windows, \n a serial driver will most likely need to be downloaded. \n Drivers for FTDI chip powered adapters can be found [here](https://ftdichip.com/drivers/vcp-drivers/)")
//	content := widget.NewRichTextFromMarkdown( " ## My serial cable is not showing up on Windows \n In order to use this application with an usb-to-rs232 adapter on windows, \n a serial driver will most likely need to be downloaded. \n Drivers for FTDI chip powered adapters can be found [here](https://ftdichip.com/drivers/vcp-drivers/)")
	return container.New(layout.NewVBoxLayout(), title, sub_title_1, content_1)

}
