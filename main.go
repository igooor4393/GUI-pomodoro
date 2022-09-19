package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"os"
	"strconv"
	"time"
)

func main() {

	//Тест функцииТест функцииТест функцииТест функцииТест функцииТест функцииТест функции
	entery := widget.NewEntry()

	//Тест функцииТест функцииТест функцииТест функцииТест функцииТест функцииТест функции

	workTime := "Work time: 00:00:07"
	restTime := "Rest time: 00:00:05"
	var cicle, cicleRest int

	musicStartProgram()
	//f, _ := os.Open("uwu-voice.mp3")
	//streamer, format, _ := mp3.Decode(f)
	//speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	//speaker.Play(streamer)

	running := false

	seconds := 0
	secondsRest := 0

	clock := widget.NewLabel("Work time: 00:00:00")
	clock.Alignment = fyne.TextAlignCenter

	clockRest := widget.NewLabel("Rest time: 00:00:00")
	clockRest.Alignment = fyne.TextAlignCenter

	ciclew := widget.NewLabel("Cycles completed: 0")

	a := app.New()
	w := a.NewWindow("PomodorGo")
	w.Resize(fyne.NewSize(400, 300))

	var btnTimeStart, btnTimeStop *widget.Button
	//btnTimeStop.Disabled()
	btnTimeStop = widget.NewButton("Start new circle", func() {

		buttonPressing() //*ЗВУК

		btnTimeStop.Disable()
		// Quit goroutine
		running = false
		seconds = 0
		secondsRest = 0
		clock.SetText("Work Time: 00:00:00")
		clockRest.SetText("Break Time: 00:00:00")
		btnTimeStart.Enable()

	})
	//btnTimeStop.Disable()
	btnTimeStart = widget.NewButton("Start work", func() {

		buttonPressing() //*ЗВУК

		btnTimeStop.Disable()
		running = !running
		go func() {

			for range time.Tick(time.Second) {
				if running {
					seconds++
					clock.SetText(formatDuration(seconds))

					btnTimeStart.Disable()

					if formatDuration(seconds) == workTime {
						btnTimeStart.SetText("Start break")
						btnTimeStart.Enable()

						fmt.Println("ЭВРИКА")
						f, _ := os.Open("uwu-voice.mp3")
						streamer, format, _ := mp3.Decode(f)
						speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
						speaker.Play(streamer)
						cicle++
						ciclew.SetText(fmt.Sprintf("Cycles completed: %d", cicle))
						break
					}
				} else if cicleRest < cicle {

					btnTimeStart.SetText("Start work")
					secondsRest++

					clockRest.SetText(formatDurationRest(secondsRest))
					btnTimeStart.Disable()

					if formatDurationRest(secondsRest) == restTime {
						btnTimeStop.Enable()
						fmt.Println("ЭВРИКА2")
						musicRestBells()
						break
					}
				} else {
					return
				}

			}
		}()

	})

	//menu
	fileMenuitem1 := fyne.NewMenuItem("Введите время", func() {

	})
	fileMenu := fyne.NewMenu("Настройки", fileMenuitem1)

	mainMenu := fyne.NewMainMenu(fileMenu)
	w.SetMainMenu(mainMenu)

	w.SetContent(container.NewVBox(
		entery,
		clock,
		clockRest,
		//label,
		//entry,

		//label2,
		//entry2,

		//buttn,
		btnTimeStart,
		btnTimeStop,
		//answer,
		ciclew,
	))

	w.ShowAndRun()
}

func formatDuration(seconds int) string {
	duration, _ := time.ParseDuration(strconv.Itoa(seconds) + "s")
	return fmt.Sprintf("Work time: %02d:%02d:%02d", int64(duration.Hours())%24, int64(duration.Minutes())%60, int64(duration.Seconds())%60)
}
func formatDurationRest(seconds int) string {
	durationRest, _ := time.ParseDuration(strconv.Itoa(seconds) + "s")
	return fmt.Sprintf("Rest time: %02d:%02d:%02d", int64(durationRest.Hours())%24, int64(durationRest.Minutes())%60, int64(durationRest.Seconds())%60)
}

//	func changeTime(app fyne.App) {
//		changeTimeWindow := app.NewWindow("Введите рабочее время")
//
//		entry := widget.NewEntry()
//		entry.Validator = validation.NewRegexp(`^[0-9]+\.?[0-9]{0,3}$`, "Not valid hourly rate")
//
// }
func musicStartProgram() {
	f, _ := os.Open("uwu-voice.mp3")

	streamer, format, _ := mp3.Decode(f)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
}
func musicRestBells() {
	f, _ := os.Open("work.mp3")

	streamer, format, _ := mp3.Decode(f)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)

}
func buttonPressing() {
	file, _ := os.Open("button.mp3")

	streamer, format, _ := mp3.Decode(file)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)

}
