package itbit

import (
	"net/http"
)

const (
	Bitcoin         = "XBT"
	Etherium        = "ETH"
	Euro            = "EUR"
	SingaporeDollar = "SGD"
	USDollar        = "USD"

	BitcoinUSDollar         = "XBTUSD"
	BitcoinSingaporeDollar  = "XBTSGD"
	BitcoinEuro             = "XBTEUR"
	EtheriumUSDollar        = "ETHUSD"
	EtheriumEuro            = "ETHEUR"
	EtheriumSingaporeDollar = "ETHSGD"
)

var (
	endpoint = "https://api.itbit.com/v1"
)

type Client struct {
	*MarketService
	*TradingService
}

func NewClient(key, secret string) *Client {
	client := &http.Client{}
	return &Client{
		newMarketService(client),
		newTradingService(client, key, secret),
	}
}
