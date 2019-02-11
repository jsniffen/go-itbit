package main

import (
	"fmt"
	"log"
	"time"

	"github.com/juliansniff/go-itbit/itbit"
)

var (
	ticker = time.NewTicker(5 * time.Second)
	client = itbit.NewClient()
)

func main() {
	for {
		<-ticker.C
		t, resp, err := client.MarketService.GetTicker(itbit.BitcoinUSDollar)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(resp.Status)
		fmt.Printf("LastPrice: %10f\n", t.LastPrice)
	}
}
