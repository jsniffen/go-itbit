package itbit

type OrderBook struct {
	Asks [][]ItBitFloat `json:"asks"`
	Bids [][]ItBitFloat `json:"bids"`
}
