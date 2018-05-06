package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

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

	endpoint := fmt.Sprintf("https://api.iextrading.com/1.0/stock/market/batch?symbols=%s&types=price", strings.Join(symbols, ","))

	if err := termbox.Init(); err != nil {
		log.Panicln(err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	quotesQueue := make(chan Quotes)
	go func() {
		for {
			start := time.Now()
			if quotes, err := GetQuotes(endpoint); err != nil {
				log.Panicln(err)
			} else {
				quotesQueue <- quotes
			}
			// Ensure we're a good consumer and only hit the API once per second.
			time.Sleep(time.Second - time.Since(start))
		}
	}()

	for {
		select {
		case quotes := <-quotesQueue:
			RenderQuotes(symbols, quotes)
		case event := <-eventQueue:
			if event.Type == termbox.EventKey && event.Key == termbox.KeyEsc {
				os.Exit(0)
			}
		}
	}
}
