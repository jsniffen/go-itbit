package itbit

import (
	"fmt"
	"net/http"
	"time"
)

type Trades struct {
	Count        int `json:"count"`
	RecentTrades []struct {
		Timestamp   time.Time `json:"timestamp"`
		MatchNumber string    `json:"matchNumber"`
		Price       float64   `json:"price,string"`
		Amount      float64   `json:"amount,string"`
	} `json:"recentTrades"`
}

// MarketService returns recent trades for the specified market
//
// since is an optional parameter
func (c *Client) GetRecentTrades(tickerSymbol, since string) (Trades, error) {
	var trades Trades

	if tickerSymbol == "" {
		return trades, fmt.Errorf("tickerSymbol is a required field, got: %s", tickerSymbol)
	}

	url := fmt.Sprintf("%s/markets/%s/trades", Endpoint, tickerSymbol)

	if since != "" {
		url = fmt.Sprintf("%s?since=%s", url, since)
	}

	err := c.doRequest(http.MethodGet, url, nil, &trades)
	if err != nil {
		return trades, fmt.Errorf("could not do request: %v", err)
	}

	return trades, nil
}
