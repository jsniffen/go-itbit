package itbit

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Order struct {
	ID                         string      `json:"id"`
	WalletID                   string      `json:"walletId"`
	Side                       string      `json:"side"`
	Instrument                 string      `json:"instrument"`
	Type                       string      `json:"type"`
	Currency                   string      `json:"currency"`
	Amount                     float64     `json:"amount,string"`
	Price                      float64     `json:"price,string"`
	AmountFilled               float64     `json:"amountFilled,string"`
	VolumeWeightedAveragePrice string      `json:"volumeWeightedAveragePrice"`
	CreatedTime                time.Time   `json:"createdTime"`
	Status                     string      `json:"status"`
	Metadata                   interface{} `json:"metadata"`
	ClientOrderIdentifier      interface{} `json:"clientOrderIdentifier"`
}

// GetOrder returns a particular order based off of a wallet ID and an order ID.
//
// Both walletID and orderID are required parameters.
func (c *Client) GetOrder(walletID, orderID string) (Order, error) {
	var order Order

	URL := fmt.Sprintf("%s/wallets/%s/orders/%s", Endpoint, walletID, orderID)

	err := c.doAuthenticatedRequest(http.MethodGet, URL, nil, &order)
	if err != nil {
		return order, fmt.Errorf("could not do authenticated GetOrder request: %v", err)
	}

	return order, nil
}

func (c *Client) GetOrders(walletID, instrument, page, perPage, status string) ([]Order, error) {
	var orders []Order

	URL := fmt.Sprintf("%s/wallets/%s/orders", Endpoint, walletID)

	p := url.Values{}

	if page != "" && perPage != "" {
		p.Set("page", page)
		p.Set("page", perPage)
	}

	if instrument != "" {
		p.Set("instrument", instrument)
	}

	if status != "" {
		p.Set("status", status)
	}

	if len(p) > 0 {
		URL = fmt.Sprintf("%s?%s", URL, p.Encode())
	}

	err := c.doAuthenticatedRequest(http.MethodGet, URL, nil, &orders)
	if err != nil {
		return orders, fmt.Errorf("could not do authenticated request: %v", err)
	}

	return orders, nil

}
