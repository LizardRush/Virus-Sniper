package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

type List struct {
	widget.BaseWidget

	// Length is a callback for returning the number of items in the list.
	Length func() int `json:"-"`

	// CreateItem is a callback invoked to create a new widget to render
	// a row in the list.
	CreateItem func() fyne.CanvasObject `json:"-"`

	// UpdateItem is a callback invoked to update a list row widget
	// to display a new row in the list. The UpdateItem callback should
	// only update the given item, it should not invoke APIs that would
	// change other properties of the list itself.
	UpdateItem func(id widget.ListItemID, item fyne.CanvasObject) `json:"-"`

	// OnSelected is a callback to be notified when a given item
	// in the list has been selected.
	OnSelected func(id widget.ListItemID) `json:"-"`

	// OnSelected is a callback to be notified when a given item
	// in the list has been unselected.
	OnUnselected func(id widget.ListItemID) `json:"-"`

	// HideSeparators hides the separators between list rows
	//
	// Since: 2.5
	HideSeparators bool
}

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
        w.SetFullScreen(true)

        container.NewAppTabs(
            container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
            container.NewTabItem("Tab 2", widget.NewLabel("World!")),
        )

        w.ShowAndRun()
        }
}
