package itbit

import (
	"net/http"

	"github.com/juliansniff/go-itbit/itbit/market"
	"github.com/juliansniff/go-itbit/itbit/trading"
)

const (
	Endpoint = "https://api.itbit.com/v1/"

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
	MarketService  *market.Service
	TradingService *trading.Service
}

func NewClient(key, secret string) *Client {
	client := &http.Client{}
	return &Client{
		MarketService:  market.NewService(client, Endpoint),
		TradingService: trading.NewService(client, Endpoint, key, secret),
	}
}
