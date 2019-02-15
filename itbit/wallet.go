package itbit

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
