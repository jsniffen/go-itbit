package itbit

import (
	"fmt"
	"net/http"
	"time"
)

type Ticker struct {
	Pair          string    `json:"pair"`
	Bid           float64   `json:"bid,string"`
	BidAmt        float64   `json:"bidAmt,string"`
	Ask           float64   `json:"ask,string"`
	AskAmt        float64   `json:"askAmt,string"`
	LastPrice     float64   `json:"lastPrice,string"`
	LastAmt       float64   `json:"lastAmt,string"`
	Volume24H     float64   `json:"volume24h,string"`
	VolumeToday   float64   `json:"volumeToday,string"`
	High24H       float64   `json:"high24h,string"`
	Low24H        float64   `json:"low24h,string"`
	HighToday     float64   `json:"highToday,string"`
	LowToday      float64   `json:"lowToday,string"`
	OpenToday     float64   `json:"openToday,string"`
	VwapToday     float64   `json:"vwapToday,string"`
	Vwap24H       float64   `json:"vwap24h,string"`
	ServerTimeUTC time.Time `json:"serverTimeUTC"`
}

// GetTicker returns Ticker for the specified market.
func (c *Client) GetTicker(tickerSymbol string) (Ticker, error) {
	var ticker Ticker

	if tickerSymbol == "" {
		return ticker, fmt.Errorf("tickerSymbol is a required field, got empty string")
	}

	url := fmt.Sprintf("%s/markets/%s/ticker", Endpoint, tickerSymbol)

	err := c.doRequest(http.MethodGet, url, nil, &ticker)
	if err != nil {
		return ticker, fmt.Errorf("could not do request: %v", err)
	}

	return ticker, nil
}
