package market

import "time"

type RecentTradesResponse struct {
	Count        int `json:"count"`
	RecentTrades []struct {
		Timestamp   time.Time `json:"timestamp"`
		MatchNumber string    `json:"matchNumber"`
		Price       float64   `json:"price,string"`
		Amount      float64   `json:"amount,string"`
	} `json:"recentTrades"`
}
