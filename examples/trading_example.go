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
		log.Panic(err)
	}
	fmt.Printf("%v", wallets)
}
