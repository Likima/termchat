package main

import (
	"fmt"
	"net"
	"time"

	ui "github.com/gizak/termui/v3"
)

func render(e ui.Event, app *App, conn net.Conn) int {
	switch e.ID {
	case "<C-c>":
		clearTerminal()
		return -1
	case "<Resize>":
		updateWidgetSizes(app)
	case "<Enter>":
		if app.message == "" {
			break
		}
		app.history = "You: " + app.message
		currentTime := time.Now().Format("15:04:05")
		app.message = currentTime + " " + app.message
		_, err := fmt.Fprintf(conn, "%s\n", app.message)
		if err != nil {
			app.history = "LOST CONNECTION"
			return -1
		}
		app.chat.Rows = append(app.chat.Rows, app.history)
		app.cursorPos = 0
		app.message = ""
		app.typing.Text = "Enter a message "
	case "<C-<Backspace>>":
		if len(app.message) > 0 && app.cursorPos > 0 {
			app.message = app.message[:app.cursorPos-1] + app.message[app.cursorPos:]
			app.cursorPos--
		}
	case "<Space>":
		app.message = app.message[:app.cursorPos] + " " + app.message[app.cursorPos:]
		app.cursorPos++
	case "<MouseLeft>":
	case "<MouseRelease>": //still dont know what to do with these.
	case "<MouseRight>":
		break
	case "<Left>":
		if app.cursorPos > 0 {
			app.cursorPos--
		}
	case "<Right>":
		if app.cursorPos < len(app.message) {
			app.cursorPos++
		}
	case "<MouseWheelDown>":
		if len(app.chat.Rows) > 0 {
			app.chat.ScrollHalfPageDown() // Scroll the page down
		}
	case "<MouseWheelUp>":
		if len(app.chat.Rows) > 0 {
			app.chat.ScrollHalfPageUp() // Scroll the page up
		}
	default:
		if len(e.ID) == 1 {
			app.message = app.message[:app.cursorPos] + e.ID + app.message[app.cursorPos:]
			app.cursorPos++
		}
	}
	app.typing.Text = app.message
	if len(app.chat.Rows) > 0 {
		app.chat.ScrollDown()
	}
	updateInputBoxText(app.typing, app.cursorPos)
	ui.Render(app.chat, app.typing)
	return 0
}
