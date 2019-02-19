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

func ExampleClient_CreateNewOrder() {
	c := itbit.NewClient("key", "secret")

	order, err := c.CreateNewOrder("walletID", itbit.Order{
		Side:                  "buy",
		Type:                  "limit",
		Currency:              itbit.Bitcoin,
		Amount:                100.0350,
		Display:               100.0350,
		Price:                 750.53,
		Instrument:            itbit.BitcoinUSDollar,
		ClientOrderIdentifier: "optional",
		PostOnly:              true,
	})
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Successfully created order: %s %.2f %s", order.Side, order.Amount, order.Currency)
	// Output: Successfully created order: buy 100.03 XBT
}

func ExampleClient_CancelOrder() {
	c := itbit.NewClient("key", "secret")

	ok, err := c.CancelOrder("walletID", "orderID")
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("%t", ok)
	// Output: true
}
