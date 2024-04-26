package main

import (
	"os"
	"os/exec"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func updateWidgetSizes(app *App) {
	termWidth, termHeight := ui.TerminalDimensions()
	app.chat.SetRect(0, 0, termWidth, (termHeight*8)/9)
	app.typing.SetRect(0, (termHeight*8)/9, termWidth, termHeight)
}

func clearTerminal() {
	cmd := exec.Command("cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func updateInputBoxText(inputBox *widgets.Paragraph, cursorPos int) {
	text := inputBox.Text
	if cursorPos < len(text) {
		text = text[:cursorPos] + "[" + text[cursorPos:cursorPos+1] + "](bg:white)" + text[cursorPos+1:] // Highlight the area with red foreground and white background
	} else {
		text += "[ ](bg:white)" // Show cursor at the end if cursorPos is out of bounds
	}
	inputBox.Text = text
}
