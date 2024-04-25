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
		ui.Render(chat, typing)
		switch e.ID {
		case "<C-c>":
			clearTerminal()
			return
		case "<Resize>":
			updateWidgetSizes()
			ui.Render(chat, typing)
			continue
		default:
			message += e.ID
			for {
				e2 := <-uiEvents
				if e2.ID == "<Enter>" {
					history += message
					history += "\n"
					message = ""
					typing.Text = "Enter a Message"
					ui.Render(chat, typing)
					break
				} else if e2.ID == "<C-<Backspace>>" {
					if len(message) == 0 {
						break
					}
					message = message[:len(message)-1]

				} else if e2.ID == "<Space>" {
					message = message + " "
				} else if e2.ID == "<Resize>" {
					updateWidgetSizes()
					ui.Render(chat, typing)
				} else {
					message += e2.ID
				}
				typing.Text = message
				ui.Render(chat, typing)
			}
		}
		updateWidgetSizes()
		chat.Text = history
		ui.Render(chat, typing)
	}
}

func clearTerminal() {
	cmd := exec.Command("cls") // Use "cls" instead of "clear" on Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}
