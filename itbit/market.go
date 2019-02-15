package itbit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type MarketService struct {
	httpClient *http.Client
	endpoint   string
}

// newMarketService returns a pointer to a MarketService.
func newMarketService(c *http.Client) *MarketService {
	return &MarketService{
		httpClient: c,
		endpoint:   endpoint + "/markets",
	}
}

// GetTicker returns TickerInfo for the specified market.
func (s *MarketService) GetTicker(tickerSymbol string) (TickerInfo, error) {
	var tickerInfo TickerInfo
	if tickerSymbol == "" {
		return tickerInfo, fmt.Errorf("tickerSymbol is a required field, got empty string")
	}
	URL := s.endpoint + "/" + tickerSymbol + "/ticker"
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return tickerInfo, err
	}
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return tickerInfo, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return tickerInfo, err
	}
	err = json.Unmarshal(b, &tickerInfo)
	if err != nil {
		return tickerInfo, err
	}
	return tickerInfo, nil
}

// GetOrderBook returns the full order book for the specified market.
func (s *MarketService) GetOrderBook(tickerSymbol string) (OrderBook, error) {
	var orderBook OrderBook
	if tickerSymbol == "" {
		return orderBook, fmt.Errorf("tickerSymbol is a required field, got empty string")
	}
	URL := s.endpoint + "/" + tickerSymbol + "/order_book"
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return orderBook, err
	}
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return orderBook, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return orderBook, err
	}
	err = json.Unmarshal(b, &orderBook)
	if err != nil {
		return orderBook, err
	}
	return orderBook, nil
}

// MarketService returns recent trades for the specified market
//
// since is an optional parameter
func (s *MarketService) GetRecentTrades(tickerSymbol, since string) (RecentTradesResponse, error) {
	var recentTrades RecentTradesResponse
	if tickerSymbol == "" {
		return recentTrades, fmt.Errorf("tickerSymbol is a required field, got: %s", tickerSymbol)
	}
	URL := s.endpoint + "/" + tickerSymbol + "/trades"
	if since != "" {
		URL = fmt.Sprintf("%s?since=%s", URL, since)
	}
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return recentTrades, err
	}
	resp, err := s.httpClient.Do(req)
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
