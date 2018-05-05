package main

import (
	"fmt"

	"github.com/maxmclau/gput"
)

// RenderQuotes will render the quotes map in the order of the provided symbols.
func RenderQuotes(symbols []string, quotes Quotes) {
	gput.Civis()
	for _, symbol := range symbols {
		fmt.Printf("%4s $%.2f\n", symbol, quotes[symbol].Price)
	}
	gput.Cuu(len(symbols))
}
