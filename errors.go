package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

//displays an error to the user
func DisplayError(w *fyne.Window, err string) {
	popup := widget.NewRichTextFromMarkdown("# " + err)
	widget.ShowPopUp(popup, (*w).Canvas())
}
