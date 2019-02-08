package itbit

import "net/http"

type Client struct {
	MarketData *MarketDataService
}

func NewClient() *Client {
	client := &http.Client{}
	return &Client{
		MarketData: newMarketDataService(client),
	}
}
