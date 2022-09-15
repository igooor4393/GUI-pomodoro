package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"time"
)

func main() {

	seconds := 0

	clock := widget.NewLabel("Time: 00:00:00")
	clock.Alignment = fyne.TextAlignCenter

	a := app.New()
	w := a.NewWindow("Недокалькулятор GUI golang")
	w.Resize(fyne.NewSize(400, 300))

	label := widget.NewLabel("enter number")
	entry := widget.NewEntry()

	label2 := widget.NewLabel("enter number")
	entry2 := widget.NewEntry()

	answer := widget.NewLabel("")

	buttn := widget.NewButton("calculate", func() {
		data1 := entry.Text
		data2 := entry2.Text
		fmt.Println(data1, data2)

		numb1, er := strconv.ParseFloat(data1, 64)
		numb2, err := strconv.ParseFloat(data2, 64)
		if err != nil || er != nil {
			answer.SetText("Ошибка ввода")
		} else {
			sum := numb1 + numb2
			res := numb1 - numb2
			mul := numb1 * numb2
			div := numb1 / numb2

			answer.SetText(fmt.Sprintf("(+) %f\n (-) %f\n (*) %f\n (/) %f\n", sum, res, mul, div))

		}
	})

	btnTime := widget.NewButton("start", func() {
		for range time.Tick(time.Second) {
			seconds++
			clock.SetText(formatDuration(seconds))
		}
	})

	w.SetContent(container.NewVBox(

		clock,
		label,
		entry,

		label2,
		entry2,

		buttn,
		btnTime,
		answer,
	))

	w.ShowAndRun()
}

func formatDuration(seconds int) string {
	duration, _ := time.ParseDuration(strconv.Itoa(seconds) + "s")
	return fmt.Sprintf("Time: %02d:%02d:%02d", int64(duration.Hours())%24, int64(duration.Minutes())%60, int64(duration.Seconds())%60)
}
