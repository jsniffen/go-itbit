package itbit

import (
	"net/http"

	"github.com/juliansniff/go-itbit/itbit/market"
	"github.com/juliansniff/go-itbit/itbit/trading"
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

type Client struct {
	*market.MarketService
	*trading.TradingService
}

func NewClient(key, secret string) *Client {
	client := &http.Client{}
	return &Client{
		market.NewMarketService(client),
		trading.NewTradingService(client, key, secret),
	}
}
