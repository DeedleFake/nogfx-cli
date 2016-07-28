package tui

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

type InputBox struct {
	screen  *Screen
	written []rune
}

func (inputBox *InputBox) SetScreen(screen *Screen) {
	inputBox.screen = screen
}

func (inputBox *InputBox) Add(ch rune) {
	inputBox.written = append(inputBox.written, ch)
}

func (inputBox *InputBox) Remove() {
	inputBox.written = inputBox.written[:len(inputBox.written)-1]
}

func (inputBox *InputBox) Get() string {
	written := string(inputBox.written)
	inputBox.written = inputBox.written[:0]
	return written
}

func (inputBox *InputBox) Draw() {
	x := 0
	for _, c := range []rune(inputBox.written) {
		termbox.SetCell(x, inputBox.screen.Height-1, c, termbox.ColorWhite, termbox.ColorDefault)
		x += runewidth.RuneWidth(c)
	}

	termbox.SetCursor(x, inputBox.screen.Height-1)

	for i := len(inputBox.written); i < inputBox.screen.Width; i++ {
		termbox.SetCell(i, inputBox.screen.Height-1, '_', termbox.ColorWhite, termbox.ColorDefault)
	}
}
