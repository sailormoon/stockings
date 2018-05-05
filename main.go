package main

import (
	"fmt"
	"os"
	"strings"
	"time"
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

	for {
		start := time.Now()
		quotes, _ := GetQuotes(endpoint)
		RenderQuotes(symbols, quotes)
		// Ensure we're a good consumer and only hit the API once per second.
		time.Sleep(time.Second - time.Since(start))
	}
}
