package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

//program entry point
func main() {
	a := app.New()

	//create a new window and resize it
	w := a.NewWindow("Saver5")
	w.Resize(fyne.Size{800, 400})

	//initializing the configuration happens here
	model := DefaultModel()

	// create tabs for each of our windows
	tabs := container.NewAppTabs(
		container.NewTabItem("Serial Port Config", SerialSelectionMenu(&model, &w)),
		container.NewTabItem("Upload", UploadMenu(&model, &w)),
		container.NewTabItem("Download", DownloadMenu(&model, &w)),
		container.NewTabItem("Help", HelpMenu()),
		container.NewTabItem("About", AboutMenu()),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	//set the content of our window to show the tabs
	w.SetContent(tabs)

	//render the window
	w.ShowAndRun()
}
