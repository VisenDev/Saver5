package main

import (
	"fyne.io/fyne/v2/container"
   "fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type View struct {
	UploadFileDisplay *widget.TextGrid
	DownloadFileDisplay *widget.TextGrid
	ConnectionStatusDisplay *widget.ProgressBarInfinite
	
	WindowContent fyne.CanvasObject
}

func (v* View) Sync(m* Model) {
	v.UploadFileDisplay.SetText(string(m.UploadFileBuffer))
	v.DownloadFileDisplay.SetText(string(m.DownloadFileBuffer[0:m.DownloadBufLen]))
}

func CreateView(w* fyne.Window, m* Model) View {
	
	// create tabs for each of our windows
	tabs := container.NewAppTabs(
      container.NewTabItemWithIcon("Home", theme.HomeIcon(), HomeMenu()),
		container.NewTabItemWithIcon("Serial Port Config", theme.SettingsIcon(),SerialSelectionMenu(m, w)),
		container.NewTabItemWithIcon("Upload", theme.UploadIcon(), UploadMenu(m, w)),
		container.NewTabItemWithIcon("Download", theme.DownloadIcon(), DownloadMenu(m, w)),
		container.NewTabItemWithIcon("Help", theme.HelpIcon(), HelpMenu()),
		//container.NewTabItemWithIcon("About", theme.InfoIcon(), AboutMenu()),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	return View{WindowContent: tabs}
}


