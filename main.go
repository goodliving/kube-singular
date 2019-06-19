package main

import (
	// "fmt"
	"fyne.io/fyne/widget"
	"fyne.io/fyne/app"
	// "github.com/satori/go.uuid"
)


func create_uuid() {
	// u1 := uuid.Must(uuid.NewV4())
	// fmt.Printf("UUIDv4: %s\n", u1)
	return
}

func main() {
	app := app.New()

	w := app.NewWindow("Hello")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("uuid", create_uuid),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))

	w.ShowAndRun()
}