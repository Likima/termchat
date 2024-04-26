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
	cursorPos := 0

	chat := widgets.NewList()
	chat.Title = "CHAT"
	chat.WrapText = true // Enable text wrapping
	chat.SetRect(0, 0, 80, 20) // Set initial size of the chat widget

	typing := widgets.NewParagraph()
	typing.Title = "MESSAGE"
	typing.Text = ""

	updateWidgetSizes := func() {
		termWidth, termHeight := ui.TerminalDimensions()
		chat.SetRect(0, 0, termWidth, (termHeight*8)/9)
		typing.SetRect(0, (termHeight*8)/9, termWidth, termHeight)
	}

	updateWidgetSizes()
	updateInputBoxText(typing, cursorPos)
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
			if message == "" {
				continue
			}
			history = "You: " + message
			chat.Rows = append(chat.Rows, history)
			cursorPos = 0
			message = ""
			typing.Text = "Enter a Message"
		case "<C-<Backspace>>":
			if len(message) > 0 && cursorPos > 0 {
				message = message[:cursorPos-1] + message[cursorPos:]
				cursorPos--
			}
		case "<Space>":
			message = message[:cursorPos] + " " + message[cursorPos:]
			cursorPos++
		case "<MouseLeft>":
		case "<MouseRelease>":
		case "<MouseRight>":
			continue
		case "<Left>":
			if cursorPos > 0 {
				cursorPos--
			}
		case "<Right>":
			if cursorPos < len(message) {
				cursorPos++
			}
		case "<MouseWheelDown>":
			if len(chat.Rows) > 0 {
				chat.ScrollHalfPageDown() // Scroll the page down
			}
		case "<MouseWheelUp>":
			if len(chat.Rows) > 0 {
				chat.ScrollHalfPageUp() // Scroll the page up
			}
		default:
			if len(e.ID) == 1 {
				message = message[:cursorPos] + e.ID + message[cursorPos:]
				cursorPos++
			}
		}
		typing.Text = message
		if len(chat.Rows) > 0 {
			chat.ScrollDown()
		}
		updateInputBoxText(typing, cursorPos)
		ui.Render(chat, typing)
	}
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
