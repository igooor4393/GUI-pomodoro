package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	//s:=1

	a := app.New()
	w := a.NewWindow("Hello World")

	//:= widget.NewLabel(string(s)))
	lable := widget.NewLabel("Привет")
	entry := widget.NewEntry()
	btn := widget.NewButton("Push", func() {
		data := entry.Text
		fmt.Println(data)
	})

	w.SetContent(container.NewVBox(
		lable,
		entry,
		btn,
	))

	w.ShowAndRun()
}
