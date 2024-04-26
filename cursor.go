package main

import (
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type InputField struct {
	*widgets.Paragraph
	cursorPos int
}

func (f *InputField) TypeRune(r rune) {
    // Insert the rune at the cursor position
    f.Text = f.Text[:f.cursorPos] + string(r) + f.Text[f.cursorPos:]
    f.cursorPos++
}

func (f *InputField) MoveCursorLeft() {
    if f.cursorPos > 0 {
        f.cursorPos--
    }
}

func (f *InputField) MoveCursorRight() {
    if f.cursorPos < len(f.Text) {
        f.cursorPos++
    }
}

func (f *InputField) Backspace() {
    if f.cursorPos > 0 {
        // Remove the character before the cursor
        f.Text = f.Text[:f.cursorPos-1] + f.Text[f.cursorPos:]
        f.cursorPos--
    }
}

func (f *InputField) Render() {
    // Add a cursor to the display
    f.Paragraph.Text = f.Text[:f.cursorPos] + "|" + f.Text[f.cursorPos:]
    termui.Render(f)
}