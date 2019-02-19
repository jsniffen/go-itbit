package itbit_test

import (
	"fmt"
	"log"

	"github.com/juliansniff/go-itbit/itbit"
)

func ExampleClient_GetAllWallets() {
	c := itbit.NewClient("key", "secret")

	wallets, err := c.GetAllWallets("userID", 0, 0)
	if err != nil {
		log.Panic(err)
	}

	for _, wallet := range wallets {
		fmt.Printf("%s:\n", wallet.Name)
		for _, balance := range wallet.Balances {
			fmt.Printf("%.2f %s", balance.AvailableBalance, balance.Currency)
		}
		fmt.Printf("\n")
	}
	// Output:
	// Wallet 1:
	// 0.00 USD0.00 XBT0.00 EUR0.00 SGD
	// Wallet 2:
	// 75631.89 USD100100.03 XBT100000.00 EUR100000.00 SGD
}

func ExampleClient_CreateNewWallet() {
	c := itbit.NewClient("key", "secret")
	wallet, err := c.CreateNewWallet("userID", "New Wallet")
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("%s:\n", wallet.Name)
	for _, balance := range wallet.Balances {
		fmt.Printf("%.2f %s Available\n", balance.AvailableBalance, balance.Currency)
	}
	// Output:
	// New Wallet:
	// 0.00 USD Available
	// 0.00 XBT Available
	// 0.00 EUR Available
	// 0.00 SGD Available
}

func ExampleClient_GetWallet() {
	c := itbit.NewClient("key", "secret")

	wallet, err := c.GetWallet("testWalletID")
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("%s:\n", wallet.Name)
	for _, balance := range wallet.Balances {
		fmt.Printf("%.2f %s Available\n", balance.AvailableBalance, balance.Currency)
	}
	// Output:
	// Test Wallet:
	// 50000.00 USD Available
	// 100.00 XBT Available
	// 100000.00 EUR Available
	// 515440.88 SGD Available
}

func ExampleClient_GetWalletBalance() {
	c := itbit.NewClient("key", "secret")
	balance, err := c.GetWalletBalance("walletID", itbit.USDollar)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("wallet has %.2f %s", balance.TotalBalance, balance.Currency)
	// Output: wallet has 0.00 USD
}
