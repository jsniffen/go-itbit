package itbit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Wallet struct {
	ID       string `json:"id"`
	UserID   string `json:"userId"`
	Name     string `json:"name"`
	Balances []struct {
		Currency         string  `json:"currency"`
		AvailableBalance float64 `json:"availableBalance,string"`
		TotalBalance     float64 `json:"totalBalance,string"`
	} `json:"balances"`
}

func (c *Client) GetAllWallets(userID string, page, perPage int) ([]Wallet, error) {
	var wallets []Wallet

	if userID == "" {
		return wallets, fmt.Errorf("userID is required, got empty string")
	}

	URL := fmt.Sprintf("%s/wallets?userId=%s", endpoint, userID)

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
