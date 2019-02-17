package itbit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type TickerInfo struct {
	Pair          string    `json:"pair"`
	Bid           float64   `json:"bid,string"`
	BidAmt        float64   `json:"bidAmt,string"`
	Ask           float64   `json:"ask,string"`
	AskAmt        float64   `json:"askAmt,string"`
	LastPrice     float64   `json:"lastPrice,string"`
	LastAmt       float64   `json:"lastAmt,string"`
	Volume24H     float64   `json:"volume24h,string"`
	VolumeToday   float64   `json:"volumeToday,string"`
	High24H       float64   `json:"high24h,string"`
	Low24H        float64   `json:"low24h,string"`
	HighToday     float64   `json:"highToday,string"`
	LowToday      float64   `json:"lowToday,string"`
	OpenToday     float64   `json:"openToday,string"`
	VwapToday     float64   `json:"vwapToday,string"`
	Vwap24H       float64   `json:"vwap24h,string"`
	ServerTimeUTC time.Time `json:"serverTimeUTC"`
}

// GetTicker returns TickerInfo for the specified market.
func (c *Client) GetTicker(tickerSymbol string) (TickerInfo, error) {
	var tickerInfo TickerInfo
	if tickerSymbol == "" {
		return tickerInfo, fmt.Errorf("tickerSymbol is a required field, got empty string")
	}
	URL := Endpoint + "/markets/" + tickerSymbol + "/ticker"
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return tickerInfo, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return tickerInfo, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return tickerInfo, err
	}
	err = json.Unmarshal(b, &tickerInfo)
	if err != nil {
		return tickerInfo, err
	}
	return tickerInfo, nil
}
