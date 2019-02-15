package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/juliansniff/go-itbit/itbit"
)

var (
	key    = os.Getenv("ITBIT_KEY")
	secret = os.Getenv("ITBIT_SECRET")
	ticker = time.NewTicker(5 * time.Second)
	client = itbit.NewClient(key, secret)
)

func main() {
	for {
		<-ticker.C
		t, err := client.GetTicker(itbit.BitcoinUSDollar)
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("LastPrice: %10f\n", t.LastPrice)
	}
}
