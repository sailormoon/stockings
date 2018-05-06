package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

func writeLine(line string, y int) {
	for index, char := range line {
		termbox.SetCell(index, y, rune(char), termbox.ColorDefault, termbox.Attribute(0))
	}
}

// RenderQuotes will render the quotes map in the order of the provided symbols.
func RenderQuotes(symbols []string, quotes Quotes) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for index, symbol := range symbols {
		writeLine(fmt.Sprintf("%-8s %8.2f", symbol, quotes[symbol].Price), index)
	}
	termbox.Flush()
}
