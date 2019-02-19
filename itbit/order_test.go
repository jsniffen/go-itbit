package itbit_test

import (
	"fmt"
	"log"

	"github.com/juliansniff/go-itbit/itbit"
)

func ExampleClient_GetOrder() {
	c := itbit.NewClient("key", "secret")

	order, err := c.GetOrder("walletID", "orderID")
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Order:\n")
	fmt.Printf("----------\n")
	fmt.Printf("Currency: %s\n", order.Currency)
	fmt.Printf("Amount: %.2f\n", order.Amount)
	fmt.Printf("Price: %.2f\n", order.Price)
	fmt.Printf("Status: %s\n", order.Status)

	// Output:
	// Order:
	// ----------
	// Currency: XBT
	// Amount: 2.50
	// Price: 650.00
	// Status: open
}

func ExampleClient_GetOrders() {
	c := itbit.NewClient("key", "secret")

	orders, err := c.GetOrders("id", "", "", "", itbit.StatusOpen)
	if err != nil {
		log.Panic(err)
	}

	for i, order := range orders {
		fmt.Printf("Order %d:\n", i+1)
		fmt.Printf("----------\n")
		fmt.Printf("Currency: %s\n", order.Currency)
		fmt.Printf("Amount: %.2f\n", order.Amount)
		fmt.Printf("Price: %.2f\n", order.Price)
		fmt.Printf("Status: %s\n", order.Status)
	}

	// Output:
	// Order 1:
	// ----------
	// Currency: XBT
	// Amount: 2.50
	// Price: 650.00
	// Status: open
}
