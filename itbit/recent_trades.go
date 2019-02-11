package itbit

import "time"

type RecentTradesResponse struct {
	Count        int `json:"count"`
	RecentTrades []struct {
		Timestamp   time.Time  `json:"timestamp"`
		MatchNumber string     `json:"matchNumber"`
		Price       ItBitFloat `json:"price"`
		Amount      ItBitFloat `json:"amount"`
	} `json:"recentTrades"`
}
