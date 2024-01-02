package main

import (
	"fmt"
	"os"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

func main() {
	var args = os.Args[1:]
	fmt.Println("args: ", args)

	application := app.New()
	welcome := application.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	clock := widget.NewLabel("")
	updateTime(clock)

	welcome.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	timeWindow := application.NewWindow("Time")
	timeWindow.SetContent(container.NewVBox(
		clock,
	))

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
	welcome.Show()
	timeWindow.Show()
	application.Run()
}
