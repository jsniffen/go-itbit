package main

import (
	"fmt"
	"log"
	"os"

	"github.com/juliansniff/go-itbit/itbit"
)

func main() {
	key := os.Getenv("ITBIT_KEY")
	secret := os.Getenv("ITBIT_SECRET")
	userID := os.Getenv("ITBIT_USER_ID")
	client := itbit.NewClient(key, secret)
	wallets, err := client.TradingService.GetAllWallets(userID, 0, 0)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%v", wallets)
}
