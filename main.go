package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	//"fyne.io/fyne/v2/container"
   //"fyne.io/fyne/v2/theme"
)

//program entry point
func main() {
	a := app.New()

	//create a new window and resize it
	w := a.NewWindow("Saver5")
	w.Resize(fyne.Size{800, 400})

	//initializing the configuration happens here
	model := DefaultModel()
	view := CreateView(&w, &model)
	
	view.Sync(&model)
   
	//// create tabs for each of our windows
	//tabs := container.NewAppTabs(
   //   container.NewTabItemWithIcon("Home", theme.HomeIcon(), HomeMenu()),
	//	container.NewTabItemWithIcon("Serial Port Config", theme.SettingsIcon(),SerialSelectionMenu(&model, &w)),
	//	container.NewTabItemWithIcon("Upload", theme.UploadIcon(), UploadMenu(&model, &w)),
	//	container.NewTabItemWithIcon("Download", theme.DownloadIcon(), DownloadMenu(&model, &w)),
	//	container.NewTabItemWithIcon("Help", theme.HelpIcon(), HelpMenu()),
	//	//container.NewTabItemWithIcon("About", theme.InfoIcon(), AboutMenu()),
	//)

	//tabs.SetTabLocation(container.TabLocationLeading)

	//set the content of our window to show the tabs
	//w.SetContent(view.WindowContent)

	//render the window
	w.ShowAndRun()
}
