package itbit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	URL := fmt.Sprintf("%s/wallets?userId=%s", Endpoint, userID)

	if page != 0 {
		URL = fmt.Sprintf("%s?page=%d", URL, page)
	}

	if perPage != 0 {
		URL = fmt.Sprintf("%s?perPage=%d", URL, perPage)
	}

	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return wallets, fmt.Errorf("error creating request: %v", err)
	}

	err = c.signRequest(req)
	if err != nil {
		return wallets, fmt.Errorf("error signing request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return wallets, fmt.Errorf("error completeing request: %v", err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return wallets, fmt.Errorf("error reading response: %v", err)
	}

	if resp.StatusCode != 200 {
		return wallets, fmt.Errorf("error status not 200: %s", string(b))
	}

	err = json.Unmarshal(b, &wallets)
	if err != nil {
		return wallets, fmt.Errorf("error Unmarshalling response into object: %v", err)
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

	URL := Endpoint + "/wallets"
	req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(bodyJSON))
	if err != nil {
		return wallet, fmt.Errorf("error creating request: %v", err)
	}

	err = c.signRequest(req)
	if err != nil {
		return wallet, fmt.Errorf("error signing request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return wallet, fmt.Errorf("could not do request: %v", err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return wallet, fmt.Errorf("error reading response: %v", err)
	}

	if resp.StatusCode != 200 {
		return wallet, fmt.Errorf("error status not 200: %s", string(b))
	}

	fmt.Printf("%v", b)
	err = json.Unmarshal(b, &wallet)
	if err != nil {
		return wallet, fmt.Errorf("could not unmarshal response body into wallet: %v", err)
	}

	return wallet, nil
}

func (c *Client) GetWallet(walletID string) (Wallet, error) {
	var wallet Wallet

	if walletID == "" {
		return wallet, fmt.Errorf("walletID required, got empty string")
	}

	URL := fmt.Sprintf("%s/wallets/%s", Endpoint, walletID)

	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return wallet, fmt.Errorf("could not create request: %v", err)
	}

	err = c.signRequest(req)
	if err != nil {
		return wallet, fmt.Errorf("could not sign request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return wallet, fmt.Errorf("could not do request: %v", err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return wallet, fmt.Errorf("could not read response body: %v", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return wallet, fmt.Errorf("%s: %s", resp.Status, string(b))
	}

	err = json.Unmarshal(b, &wallet)
	if err != nil {
		return wallet, fmt.Errorf("error Unmarshalling response body into wallet: %v", err)
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
