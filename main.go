package main

import (
	"fmt"
	"log"
	"net"

	ui "github.com/gizak/termui/v3"
)

func main() {
	conn, err1 := net.Dial("tcp", ":8080")
	if err1 != nil {
		log.Fatalf("Error starting server: %s", err1)
	}
	defer conn.Close()
	err2 := initUI()
	defer err2()

	message := ""
	history := ""

	app := &App{
		message:   message,
		history:   history,
		cursorPos: 0,
		chat:      initChat(),
		typing:    initTyping(),
	}

	updateAll(app)

	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		if render(e, app, conn) == -1 {
			clearTerminal()
			fmt.Println("Connection Interrupted!")
			break
		}
	}
}
