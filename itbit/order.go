package itbit

import (
	"bytes"
	"encoding/json"
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
	Display                    float64     `json:"display,string"`
	Price                      float64     `json:"price,string"`
	AmountFilled               float64     `json:"amountFilled,string"`
	VolumeWeightedAveragePrice string      `json:"volumeWeightedAveragePrice"`
	CreatedTime                time.Time   `json:"createdTime"`
	Status                     string      `json:"status"`
	Metadata                   interface{} `json:"metadata"`
	ClientOrderIdentifier      interface{} `json:"clientOrderIdentifier"`
	PostOnly                   bool        `json:"postOnly"`
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

// GetOrders returns an array of orders given a wallet ID.
func (c *Client) GetOrders(walletID, instrument, page, perPage, status string) ([]Order, error) {
	var orders []Order

	URL := fmt.Sprintf("%s/wallets/%s/orders", Endpoint, walletID)

	p := url.Values{}

	if page != "" && perPage != "" {
		p.Set("page", page)
		p.Set("perPage", perPage)
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

func (c *Client) CreateNewOrder(walletID string, order Order) (Order, error) {
	var response Order

	URL := fmt.Sprintf("%s/wallets/%s/orders", Endpoint, walletID)

	body, err := json.Marshal(order)
	if err != nil {
		return response, fmt.Errorf("could not marshal Order %#v: %v", order, err)
	}

	err = c.doAuthenticatedRequest(http.MethodPost, URL, bytes.NewBuffer(body), &response)
	if err != nil {
		fmt.Errorf("could not do authenticated CreateNewOrder request: %v", err)
	}

	return response, nil
}

// CancelOrder cancels an order for the specified wallet ID and order ID.
// A successful response indicates that the request was recieved, but does
// not guarantee the order was cancelled.
func (c *Client) CancelOrder(walletID, orderID string) (bool, error) {
	URL := fmt.Sprintf("%s/wallets/%s/orders/%s", Endpoint, walletID, orderID)

	err := c.doAuthenticatedRequest(http.MethodDelete, URL, nil, nil)
	if err != nil {
		return false, fmt.Errorf("could not do authenticated CancelOrder request: %v", err)
	}

	return true, nil
}
