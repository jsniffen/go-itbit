package itbit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Balance struct {
	Currency         string  `json:"currency"`
	AvailableBalance float64 `json:"availableBalance,string"`
	TotalBalance     float64 `json:"totalBalance,string"`
}

type Funding struct {
	TotalNumberOfRecords string `json:"totalNumberOfRecords"`
	CurrentPageNumber    string `json:"currentPageNumber"`
	LatestExecutionID    string `json:"latestExecutionId"`
	RecordsPerPage       string `json:"recordsPerPage"`
	FundingHistory       []struct {
		BankName                    string `json:"bankName,omitempty"`
		WithdrawalID                int    `json:"withdrawalId,omitempty"`
		HoldingPeriodCompletionDate string `json:"holdingPeriodCompletionDate,omitempty"`
		Time                        string `json:"time"`
		Currency                    string `json:"currency"`
		TransactionType             string `json:"transactionType"`
		Amount                      string `json:"amount"`
		WalletName                  string `json:"walletName"`
		Status                      string `json:"status"`
		DestinationAddress          string `json:"destinationAddress,omitempty"`
		TxnHash                     string `json:"txnHash,omitempty"`
	} `json:"fundingHistory"`
}

type Transfer struct {
	SourceWalletID      string  `json:"sourceWalletID"`
	DestinationWalletID string  `json:"destinationWalletID"`
	Amount              float64 `json:"amount,string"`
	CurrencyCode        string  `json:"currencyCode"`
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

	err := c.doAuthenticatedRequest(http.MethodGet, URL, nil, &wallets)
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

	URL := fmt.Sprintf("%s/%s", Endpoint, "wallets")

	err = c.doAuthenticatedRequest(http.MethodPost, URL, bytes.NewBuffer(bodyJSON), &wallet)
	if err != nil {
		return wallet, fmt.Errorf("could not do authenticated request: %v", err)
	}

	return wallet, nil
}

func (c *Client) GetWallet(walletID string) (Wallet, error) {
	var wallet Wallet

	if walletID == "" {
		return wallet, fmt.Errorf("walletID required, got empty string")
	}

	URL := fmt.Sprintf("%s/wallets/%s", Endpoint, walletID)

	err := c.doAuthenticatedRequest(http.MethodGet, URL, nil, &wallet)
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

	URL := fmt.Sprintf("%s/wallets/%s/balances/%s", Endpoint, walletID, currencyCode)

	err := c.doAuthenticatedRequest(http.MethodGet, URL, nil, &balance)
	if err != nil {
		return balance, fmt.Errorf("could not do authenticated request: %v", err)
	}

	return balance, nil
}

// NewWalletTransfer transfers funds from one wallet to another.
func (c *Client) NewWalletTransfer(transfer Transfer) (Transfer, error) {
	var response Transfer

	body, err := json.Marshal(transfer)
	if err != nil {
		return response, fmt.Errorf("could not marshal body: %v", err)
	}

	URL := fmt.Sprintf("%s/%s", Endpoint, "wallet_transfers")

	err = c.doAuthenticatedRequest(http.MethodPost, URL, bytes.NewBuffer(body), &response)
	if err != nil {
		return response, fmt.Errorf("could not do authenticated request: %v", err)
	}

	return response, nil
}

func (c *Client) GetFundingHistory(walletID, lastExecutionID, page, perPage string, rangeStart, rangeEnd time.Time) (Funding, error) {
	var funding Funding

	URL := fmt.Sprintf("%s/wallets/%s/trades", Endpoint, walletID)

	p := url.Values{}

	if lastExecutionID != "" {
		p.Set("lastExecutionId", lastExecutionID)
	}

	if page != "" && perPage != "" {
		p.Set("page", page)
		p.Set("perPage", perPage)
	}

	if !rangeStart.IsZero() {
		p.Set("rangeStart", rangeStart.Format(time.RFC3339))
	}

	if !rangeEnd.IsZero() {
		p.Set("rangeEnd", rangeEnd.Format(time.RFC3339))
	}

	if len(p) > 0 {
		URL = fmt.Sprintf("%s?%s", URL, p.Encode())
	}

	err := c.doAuthenticatedRequest(http.MethodGet, URL, nil, &funding)
	if err != nil {
		return funding, fmt.Errorf("could not do authenticated get funding funding reqeust: %v", err)
	}

	return funding, nil
}
