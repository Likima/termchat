package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type App struct {
    message, history string
    cursorPos        int
    chat             *widgets.List
    typing           *widgets.Paragraph
}

func initUI() func() {
    if err := ui.Init(); err != nil {
        log.Fatalf("failed to initialize termui: %v", err)
    }
    return ui.Close
}

func initChat() *widgets.List {
    chat := widgets.NewList()
    chat.Title = "CHAT"
    chat.WrapText = true
    return chat
}

func initTyping() *widgets.Paragraph {
    typing := widgets.NewParagraph()
    typing.Title = "MESSAGE"
    typing.Text = ""
    return typing
}

func updateAll(app *App) {

	updateWidgetSizes(app)
	updateInputBoxText(app.typing, app.cursorPos)
	ui.Render(app.chat, app.typing)
}