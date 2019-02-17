package main

import (
	"fmt"
	"log"
	"os"

	"github.com/juliansniff/go-itbit/itbit"
)

var (
	key    = os.Getenv("ITBIT_KEY")
	secret = os.Getenv("ITBIT_SECRET")
	userID = os.Getenv("ITBIT_USER_ID")
)

func main() {
	client := itbit.NewClient(key, secret)
	wallets, err := client.GetAllWallets(userID, 0, 0)
	if err != nil {
		log.Printf("%v", err)
	}

	for _, wallet := range wallets {
		balance, err := client.GetWalletBalance(wallet.ID, itbit.USDollar)
		if err != nil {
			log.Printf("%v", err)
		}

		fmt.Printf("Wallet %s has balance: %f %s\n", wallet.Name, balance.TotalBalance, balance.Currency)
	}
}
