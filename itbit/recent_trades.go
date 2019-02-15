package itbit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type RecentTradesResponse struct {
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
func (c *Client) GetRecentTrades(tickerSymbol, since string) (RecentTradesResponse, error) {
	var recentTrades RecentTradesResponse
	if tickerSymbol == "" {
		return recentTrades, fmt.Errorf("tickerSymbol is a required field, got: %s", tickerSymbol)
	}
	URL := endpoint + "/markets/" + tickerSymbol + "/trades"
	if since != "" {
		URL = fmt.Sprintf("%s?since=%s", URL, since)
	}
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return recentTrades, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return recentTrades, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return recentTrades, err
	}
	err = json.Unmarshal(b, &recentTrades)
	if err != nil {
		return recentTrades, err
	}
	return recentTrades, nil
}
