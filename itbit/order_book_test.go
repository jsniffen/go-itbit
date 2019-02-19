package itbit_test

import (
	"fmt"
	"log"

	"github.com/juliansniff/go-itbit/itbit"
)

func ExampleClient_GetOrderBook() {
	c := itbit.NewClient("key", "secret")

	orderBook, err := c.GetOrderBook(itbit.BitcoinUSDollar)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Asks:\n")
	for _, ask := range orderBook.Asks {
		fmt.Printf("[%.2f, %.2f]\n", ask[0], ask[1])
	}
	fmt.Printf("Bids:\n")
	for _, bid := range orderBook.Bids {
		fmt.Printf("[%.2f, %.2f]\n", bid[0], bid[1])
	}

	// Output:
	// Asks:
	// [219.82, 2.19]
	// [219.83, 6.05]
	// [220.19, 17.59]
	// [220.52, 3.36]
	// [220.53, 33.46]
	// Bids:
	// [219.40, 17.46]
	// [219.13, 53.93]
	// [219.08, 2.20]
	// [218.58, 98.73]
	// [218.20, 3.37]
}
