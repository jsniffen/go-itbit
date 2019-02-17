package itbit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Balance struct {
	Currency         string  `json:"currency"`
	AvailableBalance float64 `json:"availableBalance,string"`
	TotalBalance     float64 `json:"totalBalance,string"`
}

type Wallet struct {
	ID       string    `json:"id"`
	UserID   string    `json:"userId"`
	Name     string    `json:"name"`
	Balances []Balance `json:"balances"`
}

func (c *Client) GetAllWallets(userID string, page, perPage int) ([]Wallet, error) {
	var wallets []Wallet

	if userID == "" {
		return wallets, fmt.Errorf("userID is required, got empty string")
	}

	url := fmt.Sprintf("%s/wallets?userId=%s", Endpoint, userID)

	if page != 0 {
		url = fmt.Sprintf("%s?page=%d", url, page)
	}

	if perPage != 0 {
		url = fmt.Sprintf("%s?perPage=%d", url, perPage)
	}

	err := c.doAuthenticatedRequest(http.MethodGet, url, nil, &wallets)
	if err != nil {
		return wallets, fmt.Errorf("could not do authenticated request: %v", err)
	}

	return wallets, nil
}

// CreateNewWallet creates a new wallet and returns a reference to it.
func (c *Client) CreateNewWallet(userID, walletName string) (Wallet, error) {
	var wallet Wallet

	if userID == "" {
		return wallet, fmt.Errorf("userID is required, got empty string")
	}

	if walletName == "" {
		return wallet, fmt.Errorf("walletName is required, got empty string")
	}

	bodyMap := map[string]string{
		"name":   walletName,
		"userId": userID,
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return wallet, fmt.Errorf("could not marshal bodyMap: %v", err)
	}

	url := fmt.Sprintf("%s/%s", Endpoint, wallet)

	c.doAuthenticatedRequest(http.MethodPost, url, bytes.NewBuffer(bodyJSON), &wallet)

	return wallet, nil
}

func (c *Client) GetWallet(walletID string) (Wallet, error) {
	var wallet Wallet

	if walletID == "" {
		return wallet, fmt.Errorf("walletID required, got empty string")
	}

	url := fmt.Sprintf("%s/wallets/%s", Endpoint, walletID)

	err := c.doAuthenticatedRequest(http.MethodGet, url, nil, &wallet)
	if err != nil {
		fmt.Errorf("could not do authenticated request: %v", err)
	}

	return wallet, nil
}

// GetWalletBalance returns the Balance of a wallet given a Wallet ID and a valid Currency Code.
func (c *Client) GetWalletBalance(walletID, currencyCode string) (Balance, error) {
	var balance Balance

	if walletID == "" {
		return balance, fmt.Errorf("walletID is required, got empty string")
	}

	if currencyCode == "" {
		return balance, fmt.Errorf("currencyCode is required, got empty string")
	}

	url := fmt.Sprintf("%s/wallets/%s/balances/%s", Endpoint, walletID, currencyCode)

	err := c.doAuthenticatedRequest(http.MethodGet, url, nil, &balance)
	if err != nil {
		return balance, fmt.Errorf("could not do authenticated request: %v", err)
	}

	return balance, nil
}
