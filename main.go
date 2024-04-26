package main

import (
	ui "github.com/gizak/termui/v3"
)


func main() {
	err := initUI()
	defer err()

	message := ""
	history := ""

	app := &App{
		message:    message,
		history:    history,
		cursorPos:  0,
		chat:       initChat(),
		typing:     initTyping(),
	}

	updateAll(app)

	uiEvents := ui.PollEvents()

	for {
		e:= <-uiEvents
		render(e, app)
	}
}