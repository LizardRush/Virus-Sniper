package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func main() {
    a := app.New()
    for i := 0; i < 10; i++ {
		w := a.NewWindow("Hello Fyne!")

        a.Settings().SetTheme(&myTheme{})

        hello := widget.NewLabel("Hello Fyne!")
        w.SetContent(container.NewVBox(
            hello,
            widget.NewButton("Hi!", func() {
                hello.SetText("beesechurgur")
            }),
        ))
        w.Resize(fyne.NewSize(500, 100))

        w.ShowAndRun()
        }
}
