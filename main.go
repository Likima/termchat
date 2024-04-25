package main

import (
	"log"
	"os"
	"os/exec"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	message, history := "", ""

	chat := widgets.NewParagraph()
	chat.Title = "CHAT"
	chat.Text = "Ctrl-C to quit, Type to do stuff"

	typing := widgets.NewParagraph()
	typing.Title = "Message"
	typing.Text = "Enter a Message"

	updateWidgetSizes := func() {
		termWidth, termHeight := ui.TerminalDimensions()
		chat.SetRect(0, 0, termWidth, (termHeight*8)/9)
		typing.SetRect(0, (termHeight*8)/9, termWidth, termHeight)
	}

	updateWidgetSizes()
	ui.Render(chat, typing)

	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		updateWidgetSizes()
		switch e.ID {
		case "<C-c>":
			clearTerminal()
			return
		case "<Resize>":
			updateWidgetSizes()
		case "<Enter>":
			history += message + "\n"
			message = ""
			typing.Text = "Enter a Message"
		case "<C-<Backspace>>":
			if len(message) > 0 {
				message = message[:len(message)-1]
			}
		case "<Space>":
			message += " "
		default:
			message += e.ID
		}
		typing.Text = message
		chat.Text = history
		ui.Render(chat, typing)
	}
}

func clearTerminal() {
	cmd := exec.Command("cls") // Use "cls" instead of "clear" on Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}
