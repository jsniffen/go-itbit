package itbit

import (
	"fmt"
	"net/http"
)

const MarketDataEndpoint = ItBitEndpoint + "/markets/"

type MarketDataService struct {
	httpClient *http.Client
}

func newMarketDataService(client *http.Client) *MarketDataService {
	return &MarketDataService{client}
}

// GetTicker returns ticker info for the specified market.
func (s *MarketDataService) GetTicker(tickerSymbol string) (*http.Response, error) {
	if tickerSymbol == "" {
		return nil, fmt.Errorf("tickerSymbol is a required field, got: %s", tickerSymbol)
	}
	URL := MarketDataEndpoint + tickerSymbol + "/ticker"
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}
	return s.httpClient.Do(req)
}

// GetOrderBook returns the full order book for the specified market.
func (s *MarketDataService) GetOrderBook(tickerSymbol string) (*http.Response, error) {
	if tickerSymbol == "" {
		return nil, fmt.Errorf("tickerSymbol is a required field, got: %s", tickerSymbol)
	}
	URL := MarketDataEndpoint + tickerSymbol + "/order_book"
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}
	return s.httpClient.Do(req)
}

// MarketDataService returns recent trades for the specified market
//
// since is an optional parameter
func (s *MarketDataService) GetRecentTrades(tickerSymbol, since string) (*http.Response, error) {
	if tickerSymbol == "" {
		return nil, fmt.Errorf("tickerSymbol is a required field, got: %s", tickerSymbol)
	}
	URL := MarketDataEndpoint + tickerSymbol + "/trades"
	if since != "" {
		URL = fmt.Sprintf("%s?since=%s", URL, since)
	}
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}
	return s.httpClient.Do(req)
}
