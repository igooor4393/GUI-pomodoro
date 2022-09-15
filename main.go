package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func main() {

	a := app.New()
	w := a.NewWindow("Недокалькулятор GUI golang")

	//label := widget.NewLabel("Привет")
	//entry := widget.NewEntry()
	//btn := widget.NewButton("Push", func() {
	//	data := entry.Text
	//	fmt.Println(data)
	//})
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

			//sum1 := strconv.FormatFloat(sum, 'E', -1, 64)
			//res1 := strconv.FormatFloat(res, 'E', -1, 64)
			//mul1 := strconv.FormatFloat(mul, 'E', -1, 64)
			//div1 := strconv.FormatFloat(div, 'E', -1, 64)

			//answer.SetText("(+)" + sum1 + "/n" + "(-)" + res1 + "/n" + "(*)" + mul1 + "/n" + "(/)" + div1 + "/n")
			answer.SetText(fmt.Sprintf("(+) %f\n (-) %f\n (*) %f\n (/) %f\n", sum, res, mul, div))

		}
	})

	w.SetContent(container.NewVBox(
		//label,
		//entry,
		//btn,
		label,
		entry,

		label2,
		entry2,

		buttn,

		answer,
	))

	w.ShowAndRun()
}
