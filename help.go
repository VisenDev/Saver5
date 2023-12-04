package main

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// HelpMenu returns the UI for the help menu.
func HelpMenu() fyne.CanvasObject {

	driver_text := widget.NewLabel("To use this application with a USB-to-RS232 adapter on Windows,\na serial driver will most likely need to be downloaded.\nDrivers for FTDI chip-powered adapters can be found here: ")
	linkURL, _ := url.Parse("https://ftdichip.com/drivers/vcp-drivers/")
	link := widget.NewHyperlink("FTDI Drivers", linkURL)
	driver_content := container.NewVBox(driver_text, link)
	driver_accordion := widget.NewAccordionItem("No Serial Ports Found on Windows", driver_content)

	ports_text := widget.NewLabel("Your computer does not have any serial ports,\nyou need to use a USB-to-RS232 adapter and install its driver\nThen relaunch or refresh the application")
	ports_accordion := widget.NewAccordionItem("No Serial Ports Found", ports_text)

	page_content := widget.NewAccordion(driver_accordion, ports_accordion)

	return container.NewVBox(
		page_content,
	)
}
