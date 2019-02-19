package itbit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type OrderBook struct {
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
}

func (ob *OrderBook) UnmarshalJSON(b []byte) error {
	var StringOrderBook struct {
		Asks [][]string `json:"asks"`
		Bids [][]string `json:"bids"`
	}

	err := json.Unmarshal(b, &StringOrderBook)
	if err != nil {
		return err
	}

	asks := [][]float64{}
	for i := range StringOrderBook.Asks {
		ask := []float64{}
		for j := range StringOrderBook.Asks[i] {
			f, err := strconv.ParseFloat(StringOrderBook.Asks[i][j], 64)
			if err != nil {
				return err
			}
			ask = append(ask, f)
		}
		asks = append(asks, ask)
	}
	bids := [][]float64{}
	for i := range StringOrderBook.Bids {
		bid := []float64{}
		for j := range StringOrderBook.Bids[i] {
			f, err := strconv.ParseFloat(StringOrderBook.Bids[i][j], 64)
			if err != nil {
				return err
			}
			bid = append(bid, f)
		}
		bids = append(bids, bid)
	}

	ob.Asks = asks
	ob.Bids = bids
	return nil
}

// GetOrderBook returns the full order book for the specified market.
func (c *Client) GetOrderBook(tickerSymbol string) (OrderBook, error) {
	var orderBook OrderBook

	if tickerSymbol == "" {
		return orderBook, fmt.Errorf("tickerSymbol is a required field, got empty string")
	}

	url := fmt.Sprintf("%s/markets/%s/order_book", Endpoint, tickerSymbol)

	err := c.doRequest(http.MethodGet, url, nil, &orderBook)
	if err != nil {
		return orderBook, fmt.Errorf("could not do request: %v", err)
	}

	return orderBook, nil
}
