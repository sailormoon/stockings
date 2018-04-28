package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/maxmclau/gput"
)

const Endpoint = "https://api.iextrading.com/1.0/stock/market/batch?symbols=%s&types=price"

type Quote struct {
	Price float64 `json:"price"`
}

func GetQuotes(endpoint string) (map[string]Quote, error) {
	response, err := http.Get(endpoint)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	var tickers map[string]Quote
	return tickers, json.NewDecoder(response.Body).Decode(&tickers)
}

func RenderQuotes(symbols []string, quotes map[string]Quote) {
	for _, symbol := range symbols {
		fmt.Printf("%4s $%.2f\n", symbol, quotes[symbol].Price)
	}
	gput.Cuu(len(symbols))
}

func main() {
	symbols := os.Args[1:]
	if len(symbols) == 0 {
		fmt.Println("./stockings [tickers]")
		os.Exit(1)
	}

	for i, symbol := range symbols {
		symbols[i] = strings.ToUpper(symbol)
	}
	endpoint := fmt.Sprintf(Endpoint, strings.Join(symbols, ","))

	gput.Civis()
	for {
		start := time.Now()
		quotes, _ := GetQuotes(endpoint)
		RenderQuotes(symbols, quotes)
		time.Sleep(time.Second - time.Since(start))
	}
}
