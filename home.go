package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/canvas"
	"strings"
)

import _ "embed"
//go:embed img/kager-industries.png
var embedded_logo string

// Home returns the UI for the home menu.
func HomeMenu() fyne.CanvasObject {
   title := widget.NewLabelWithStyle("Welcome to Saver5!", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	
	//logo := canvas.NewImageFromFile("kager-industries.png")
	logo := canvas.NewImageFromReader(strings.NewReader(embedded_logo), "kager-industries.png")
	logo.SetMinSize(fyne.NewSize(100 * logo.Aspect(), 100))
	logo_wrapper := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), logo, layout.NewSpacer())
  
   page_content := widget.NewRichTextFromMarkdown("For **updates** of this application check our github page [link](https://github.com/VisenDev/Saver5)")
   page_content2 := widget.NewRichTextFromMarkdown("\n Written by: \n - Robert Burnett \n \n Designed by:\n - Robert Burnett\n - Connor Bumann\n - Elizabeth Salyards\n - Eduardo Palacios")
   return container.New(layout.NewVBoxLayout(), logo_wrapper, title, page_content, page_content2)
}
