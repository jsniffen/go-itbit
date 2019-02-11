package market

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Service struct {
	httpClient *http.Client
	endpoint   string
}

// NewService returns a pointer to a Service.
func NewService(c *http.Client, baseEndpoint string) *Service {
	return &Service{
		httpClient: c,
		endpoint:   baseEndpoint + "/markets/",
	}
}

// GetTicker returns TickerInfo for the specified market.
func (s *Service) GetTicker(tickerSymbol string) (TickerInfo, *http.Response, error) {
	var tickerInfo TickerInfo
	if tickerSymbol == "" {
		return tickerInfo, nil, fmt.Errorf("tickerSymbol is a required field, got empty string")
	}
	URL := s.endpoint + tickerSymbol + "/ticker"
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return tickerInfo, nil, err
	}
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return tickerInfo, resp, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return tickerInfo, resp, err
	}
	err = json.Unmarshal(b, &tickerInfo)
	if err != nil {
		return tickerInfo, resp, err
	}
	return tickerInfo, resp, nil
}

// GetOrderBook returns the full order book for the specified market.
func (s *Service) GetOrderBook(tickerSymbol string) (OrderBook, *http.Response, error) {
	var orderBook OrderBook
	if tickerSymbol == "" {
		return orderBook, nil, fmt.Errorf("tickerSymbol is a required field, got empty string")
	}
	URL := s.endpoint + tickerSymbol + "/order_book"
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return orderBook, nil, err
	}
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return orderBook, resp, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return orderBook, resp, err
	}
	err = json.Unmarshal(b, &orderBook)
	if err != nil {
		return orderBook, resp, err
	}
	return orderBook, resp, nil
}

// Service returns recent trades for the specified market
//
// since is an optional parameter
func (s *Service) GetRecentTrades(tickerSymbol, since string) (RecentTradesResponse, *http.Response, error) {
	var recentTrades RecentTradesResponse
	if tickerSymbol == "" {
		return recentTrades, nil, fmt.Errorf("tickerSymbol is a required field, got: %s", tickerSymbol)
	}
	URL := s.endpoint + tickerSymbol + "/trades"
	if since != "" {
		URL = fmt.Sprintf("%s?since=%s", URL, since)
	}
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return recentTrades, nil, err
	}
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return recentTrades, resp, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return recentTrades, resp, err
	}
	err = json.Unmarshal(b, &recentTrades)
	if err != nil {
		return recentTrades, resp, err
	}
	return recentTrades, resp, nil
}
