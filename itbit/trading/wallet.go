package trading

type Wallet struct {
	ID       string `json:"id"`
	UserID   string `json:"userId"`
	Name     string `json:"name"`
	Balances []struct {
		Currency         string `json:"currency"`
		AvailableBalance string `json:"availableBalance"`
		TotalBalance     string `json:"totalBalance"`
	} `json:"balances"`
}
