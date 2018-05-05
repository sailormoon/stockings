package main

import (
	"encoding/json"
	"net/http"
)

// Quote represents the last price seen for a symbol.
type Quote struct {
	Price float64 `json:"price"`
}

// Quotes is a collection of Quote structs keyed by symbol (e.g. AAPL -> Quote{175.00}).
type Quotes map[string]Quote

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
