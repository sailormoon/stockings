package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nsf/termbox-go"
)

func main() {
	symbols := os.Args[1:]
	if len(symbols) == 0 {
		fmt.Println("./stockings [tickers]")
		os.Exit(1)
	}

	for i, symbol := range symbols {
		symbols[i] = strings.ToUpper(symbol)
	}

	if err := termbox.Init(); err != nil {
		log.Panicln(err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go PollEvents(eventQueue)

	endpoint := fmt.Sprintf("https://api.iextrading.com/1.0/stock/market/batch?symbols=%s&types=price", strings.Join(symbols, ","))
	quotesQueue := make(chan Quotes)
	go PollQuotes(endpoint, quotesQueue)

	for {
		select {
		case quotes := <-quotesQueue:
			// TODO: Probably force a render every iteration if key presses are going to affect the UI.
			RenderQuotes(symbols, quotes)
		case event := <-eventQueue:
			HandleEvent(event)
		}
	}
}
