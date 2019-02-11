package itbit

import "time"

type TickerInfo struct {
	Pair          string     `json:"pair"`
	Bid           ItBitFloat `json:"bid"`
	BidAmt        ItBitFloat `json:"bidAmt"`
	Ask           ItBitFloat `json:"ask"`
	AskAmt        ItBitFloat `json:"askAmt"`
	LastPrice     ItBitFloat `json:"lastPrice"`
	LastAmt       ItBitFloat `json:"lastAmt"`
	Volume24H     ItBitFloat `json:"volume24h"`
	VolumeToday   ItBitFloat `json:"volumeToday"`
	High24H       ItBitFloat `json:"high24h"`
	Low24H        ItBitFloat `json:"low24h"`
	HighToday     ItBitFloat `json:"highToday"`
	LowToday      ItBitFloat `json:"lowToday"`
	OpenToday     ItBitFloat `json:"openToday"`
	VwapToday     ItBitFloat `json:"vwapToday"`
	Vwap24H       ItBitFloat `json:"vwap24h"`
	ServerTimeUTC time.Time  `json:"serverTimeUTC"`
}
