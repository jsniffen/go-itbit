package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/juliansniff/go-itbit/itbit"
)

func main() {
	client := itbit.NewClient()
	resp, err := client.MarketData.GetTicker(itbit.BitcoinUSDollar)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%s\n", string(b))
}
