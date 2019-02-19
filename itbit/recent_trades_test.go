package itbit_test

import (
	"fmt"
	"log"

	"github.com/juliansniff/go-itbit/itbit"
)

func ExampleClient_GetRecentTrades() {
	c := itbit.NewClient("key", "secret")

	trades, err := c.GetRecentTrades(itbit.BitcoinUSDollar, "")
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("%d Trades:\n", trades.Count)
	for i, trade := range trades.RecentTrades {
		fmt.Printf("%d) Price: %.2f, Amount: %.4f\n", i+1, trade.Price, trade.Amount)
	}

	// Output:
	// 3 Trades:
	// 1) Price: 351.45, Amount: 0.0001
	// 2) Price: 352.00, Amount: 0.0001
	// 3) Price: 351.45, Amount: 0.0001
}
