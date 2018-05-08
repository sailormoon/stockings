package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

func writeLine(line string, y int, attribute termbox.Attribute) {
	for index, char := range line {
		termbox.SetCell(index, y, rune(char), attribute, termbox.ColorDefault)
	}
}

func color(percentDiff float64) termbox.Attribute {
	if percentDiff >= 0 {
		return termbox.ColorGreen
	}
	return termbox.ColorRed
}

func percentageDifference(start float64, current float64) float64 {
	return (current - start) / current * 100
}

// RenderQuotes will render the quotes map in the order of the provided symbols.
func RenderQuotes(symbols []string, quotes Quotes, states States) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for index, symbol := range symbols {
		lastPrice := quotes[symbol].Price
		percentDiff := percentageDifference(states[symbol].OpenHighLowClose.Open.Price, lastPrice)
		writeLine(fmt.Sprintf("%-8s %8.2f %6.2f%%", symbol, lastPrice,
			percentDiff), index, color(percentDiff))
	}
	termbox.Flush()
}
