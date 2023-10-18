package main

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// HelpMenu returns the UI for the help menu.
func HelpMenu() fyne.CanvasObject {

	title := widget.NewLabelWithStyle("Help", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	subTitle1 := widget.NewLabelWithStyle("My Serial Cable is Not Showing up on Windows", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})

	content1 := widget.NewLabel("To use this application with a USB-to-RS232 adapter on Windows, a serial driver will most likely need to be downloaded.")
	content2 := widget.NewLabel("Drivers for FTDI chip-powered adapters can be found here:")

	linkURL, _ := url.Parse("https://ftdichip.com/drivers/vcp-drivers/")
	link := widget.NewHyperlink("FTDI Drivers", linkURL)

	contentContainer := container.NewVBox(content1, content2, container.NewHBox(layout.NewSpacer(), link, layout.NewSpacer()))

	return container.NewVBox(
		layout.NewSpacer(),
		title,
		layout.NewSpacer(),
		subTitle1,
		layout.NewSpacer(),
		contentContainer,
		layout.NewSpacer(),
	)
}
