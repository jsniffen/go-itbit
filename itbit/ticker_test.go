package itbit_test

import (
	"fmt"
	"log"

	"github.com/juliansniff/go-itbit/itbit"
)

func ExampleClient_GetTicker() {
	c := itbit.NewClient("key", "secret")

	ticker, err := c.GetTicker(itbit.BitcoinUSDollar)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("%s:\n", ticker.Pair)
	fmt.Printf("Last Price: %.2f\n", ticker.LastPrice)
	fmt.Printf("Last Amount: %.4f\n", ticker.LastAmt)

	// Output:
	// XBTUSD:
	// Last Price: 618.00
	// Last Amount: 0.0004
}
