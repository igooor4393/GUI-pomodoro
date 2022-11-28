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
	var btnTimeStart, btnTimeStop *widget.Button
	var cicle, cicleRest int

	ic, _ := fyne.LoadResourceFromPath("Pomodor.jpg")

	workTime := "Work time: 00:00:07"
	restTime := "Rest time: 00:00:05"

	selWorkTime := widget.NewSelect([]string{"Work time: 00:00:05", "Work time: 00:15:00", "Work time: 00:20:00", "Work time: 00:25:00"}, func(s string) {
		workTime = s

	})
	selWorkTime.PlaceHolder = "Select work time:"

	selRestTime := widget.NewSelect([]string{"Rest time: 00:00:03", "Rest time: 00:03:00", "Rest time: 00:05:00", "Rest time: 00:10:00"}, func(s string) {
		restTime = s
	})
	selRestTime.PlaceHolder = "Select rest time:"

	musicStartProgram()

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
	w.SetIcon(ic)

	btnTimeStop = widget.NewButton("Start new circle", func() {

		buttonPressing() //*sound

		btnTimeStop.Disable()
		running = false
		seconds = 0
		secondsRest = 0
		clock.SetText("Work Time: 00:00:00")
		clockRest.SetText("Break Time: 00:00:00")
		btnTimeStart.Enable()

	})

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
						musicStartProgram()
						btnTimeStart.SetText("Start break")
						btnTimeStart.Enable()
						cicle++
						break
					}
				} else if cicleRest < cicle {

					btnTimeStart.SetText("Start work")
					secondsRest++

					clockRest.SetText(formatDurationRest(secondsRest))
					btnTimeStart.Disable()

					if formatDurationRest(secondsRest) == restTime {
						musicRestBells()
						btnTimeStop.Enable()
						ciclew.SetText(fmt.Sprintf("Cycles completed: %d", cicle))
						break
					}
				} else {
					return
				}
			}
		}()
	})

	w.SetContent(container.NewVBox(

		selWorkTime,
		selRestTime,
		clock,
		clockRest,
		btnTimeStart,
		btnTimeStop,
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

func musicStartProgram() {
	f, _ := os.Open("sounds/uwu-voice.mp3")

	streamer, format, _ := mp3.Decode(f)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
}
func musicRestBells() {
	f, _ := os.Open("sounds/work.mp3")
	streamer, format, _ := mp3.Decode(f)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)

}
func buttonPressing() {
	file, _ := os.Open("sounds/button.mp3")
	streamer, format, _ := mp3.Decode(file)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)

}
