package market

type OrderBook struct {
	Asks [][]float64 `json:"asks,string"`
	Bids [][]float64 `json:"bids,string"`
}
