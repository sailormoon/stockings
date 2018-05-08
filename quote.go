package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// Quote represents the last price seen for a symbol.
type Quote struct {
	Price float64 `json:"price"`
}

// Quotes is a collection of Quote structs keyed by symbol (e.g. AAPL -> Quote{175.00}).
type Quotes map[string]Quote

// State is the state queried at program start up using the Batch API, and then never queried again.
type State struct {
	OpenHighLowClose struct {
		Open  Quote `json:"open"`
		Close Quote `json:"close"`
	} `json:"ohlc"`
}

// States is a map of symbol -> State.
type States map[string]State

// GetStates attempts to unmarshall the initial program state.
func GetStates(symbols []string) (States, error) {
	endpoint := fmt.Sprintf("https://api.iextrading.com/1.0/stock/market/batch?symbols=%s&types=ohlc", strings.Join(symbols, ","))
	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer response.Body.Close()

	var states States
	return states, json.NewDecoder(response.Body).Decode(&states)
}

// GetQuotes hits the passed in endpoint and attempts to parse it as JSON.
func GetQuotes(endpoint string) (Quotes, error) {
	response, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var tickers map[string]Quote
	return tickers, json.NewDecoder(response.Body).Decode(&tickers)
}

// PollQuotes polls for updated quotes every second and sends them down the channel.
func PollQuotes(endpoint string, quotesQueue chan<- Quotes) {
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
}
