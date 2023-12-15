package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type View struct {
	UploadFileDisplay       *widget.TextGrid
	DownloadFileDisplay     *widget.TextGrid
	ConnectionStatusDisplay *widget.ProgressBarInfinite
	ConnectionStatusLabel   *widget.Label
	// connection_label.SetText("Not Connected")

	//WindowContent fyne.CanvasObject
}

func (v *View) Sync(m *Model) {
	v.UploadFileDisplay.SetText(string(m.UploadFileBuffer))

	if m.DownloadFileBuffer != nil {
		v.DownloadFileDisplay.SetText(string(m.DownloadFileBuffer[0:m.DownloadBufLen]))
	}

	if v.ConnectionStatusLabel != nil {
		if m.Port != nil {
			v.ConnectionStatusDisplay.Start()
			v.ConnectionStatusLabel.SetText("Connected to " + m.Config.Port)
		} else {
			v.ConnectionStatusDisplay.Stop()
			v.ConnectionStatusLabel.SetText("Not Connected")
		}
	}
}

func CreateView(w *fyne.Window, m *Model) View {
	view := View{nil, nil, nil, nil}

	// create tabs for each of our windows
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), HomeMenu()),
		container.NewTabItemWithIcon("Serial Port Config", theme.SettingsIcon(), SerialSelectionMenu(m, w, &view)),
		container.NewTabItemWithIcon("Upload", theme.UploadIcon(), UploadMenu(m, w, &view)),
		container.NewTabItemWithIcon("Download", theme.DownloadIcon(), DownloadMenu(m, w, &view)),
		container.NewTabItemWithIcon("Help", theme.HelpIcon(), HelpMenu()),
		// container.NewTabItemWithIcon("About", theme.InfoIcon(), AboutMenu()),
	)

	tabs.SetTabLocation(container.TabLocationLeading)
	(*w).SetContent(tabs)
	//view.WindowContent = tabs

	return view // View{WindowContent: tabs}
}
