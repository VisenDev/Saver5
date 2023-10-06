package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func DisplayError(w *fyne.Window, err string) {
	popup := widget.NewRichTextFromMarkdown("# " + err )
   widget.ShowPopUp(popup, (*w).Canvas())
}
