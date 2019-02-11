package itbit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const MarketDataEndpoint = ItBitEndpoint + "/markets/"

type MarketDataService struct {
	httpClient *http.Client
}

func newMarketDataService(client *http.Client) *MarketDataService {
	return &MarketDataService{client}
}

// GetTicker returns TickerInfo for the specified market.
func (s *MarketDataService) GetTicker(tickerSymbol string) (TickerInfo, *http.Response, error) {
	var tickerInfo TickerInfo
	if tickerSymbol == "" {
		return tickerInfo, nil, fmt.Errorf("tickerSymbol is a required field, got empty string")
	}
	URL := MarketDataEndpoint + tickerSymbol + "/ticker"
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
func (s *MarketDataService) GetOrderBook(tickerSymbol string) (OrderBook, *http.Response, error) {
	var orderBook OrderBook
	if tickerSymbol == "" {
		return orderBook, nil, fmt.Errorf("tickerSymbol is a required field, got empty string")
	}
	URL := MarketDataEndpoint + tickerSymbol + "/order_book"
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

// MarketDataService returns recent trades for the specified market
//
// since is an optional parameter
func (s *MarketDataService) GetRecentTrades(tickerSymbol, since string) (RecentTradesResponse, *http.Response, error) {
	var recentTrades RecentTradesResponse
	if tickerSymbol == "" {
		return recentTrades, nil, fmt.Errorf("tickerSymbol is a required field, got: %s", tickerSymbol)
	}
	URL := MarketDataEndpoint + tickerSymbol + "/trades"
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
