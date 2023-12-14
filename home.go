package main

import (
	//"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Home returns the UI for the home menu.
func HomeMenu() fyne.CanvasObject {
   title := widget.NewLabelWithStyle("Welcome to Saver5!", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
  
   page_content := widget.NewRichTextFromMarkdown("For updates of this application check our github page [link](https://github.com/VisenDev/Saver5)")
   page_content2 := widget.NewRichTextFromMarkdown("\n Written by:\n - Robert Burnett\n \n Designed by:\n - Robert Burnett\n - Connor Bumann\n - Elizabeth Salyards\n - Eduardo Palacious ")
   return container.New(layout.NewVBoxLayout(), title, page_content, page_content2)
}
