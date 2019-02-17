package itbit_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/juliansniff/go-itbit/itbit"
)

func handleGetWalletBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	currency := vars["currency"]
	response := fmt.Sprintf(`{
				"currency": "%s",
				"availableBalance": "0",
				"totalBalance": "0"
			}`, currency)
	fmt.Fprintf(w, response)
}

func handleGetWallet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	response := fmt.Sprintf(`{
			"id": "%s",
			"userId": "userID",
			"name": "Test Wallet",
			"balances": [
				{
					"currency": "USD",
					"availableBalance": "50000.0000000",
					"totalBalance": "50000.0000000"
				},
				{
					"currency": "XBT",
					"availableBalance": "100.00000000",
					"totalBalance": "100.00000000"
				},
				{
					"currency": "EUR",
					"availableBalance": "100000.00000000",
					"totalBalance": "100000.00000000"
				},
				{
					"currency": "SGD",
					"availableBalance": "515440.88288502",
					"totalBalance": "515432.74228603"
				}
			]
		}`, id)
	fmt.Fprintf(w, response)
}

func handleGetAllWallets(w http.ResponseWriter, r *http.Request) {
	response := `[
			{
				"id": "f46eb6b0-2a7f-4c07-898e-856d58568fde",
				"userId": "e9c856db-c726-46b2-ace0-a806671105fb",
				"name": "Wallet 1",
				"balances": [
					{
						"currency": "USD",
						"availableBalance": "0.00000000",
						"totalBalance": "0.00000000"
					},
					{
						"currency": "XBT",
						"availableBalance": "0.00000000",
						"totalBalance": "0.00000000"
					},
					{
						"currency": "EUR",
						"availableBalance": "0.00000000",
						"totalBalance": "0.00000000"
					},
					{
						"currency": "SGD",
						"availableBalance": "0.00000000",
						"totalBalance": "0.00000000"
					}
				]
			},
			{
				"id": "2224db00-da07-46f7-ba37-b877e1cd455a",
				"userId": "e9c856db-c726-46b2-ace0-a806671105fb",
				"name": "Wallet 2",
				"balances": [
					{
						"currency": "USD",
						"availableBalance": "75631.88785992",
						"totalBalance": "75631.88785999"
					},
					{
						"currency": "XBT",
						"availableBalance": "100100.03000000",
						"totalBalance": "100100.03000000"
					},
					{
						"currency": "EUR",
						"availableBalance": "100000.00000000",
						"totalBalance": "100000.00000000"
					},
					{
						"currency": "SGD",
						"availableBalance": "100000.00000000",
						"totalBalance": "100000.00000000"
					}
				]
			}
		]`
	fmt.Fprintf(w, response)
}

func handleCreateNewWallet(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	var body map[string]string
	_ = json.Unmarshal(b, &body)

	name := body["name"]
	userID := body["userId"]

	response := fmt.Sprintf(`{
				"id": "walletID",
				"userId": "%s",
				"name": "%s",
				"balances": [
					{
						"currency": "USD",
						"availableBalance": "0.00000000",
						"totalBalance": "0.00000000"
					},
					{
						"currency": "XBT",
						"availableBalance": "0.00000000",
						"totalBalance": "0.00000000"
					},
					{
						"currency": "EUR",
						"availableBalance": "0.00000000",
						"totalBalance": "0.00000000"
					},
					{
						"currency": "SGD",
						"availableBalance": "0.00000000",
						"totalBalance": "0.00000000"
					}
				]
			}`, userID, name)
	fmt.Fprintf(w, response)
}

func TestMain(m *testing.M) {
	r := mux.NewRouter()
	r.HandleFunc("/wallets/{id}/balances/{currency}", handleGetWalletBalance).Methods(http.MethodGet)
	r.HandleFunc("/wallets/{id}", handleGetWallet).Methods(http.MethodGet)
	r.HandleFunc("/wallets", handleGetAllWallets).Methods(http.MethodGet)
	r.HandleFunc("/wallets", handleCreateNewWallet).Methods(http.MethodPost)
	ts := httptest.NewServer(r)
	defer ts.Close()

	itbit.Endpoint = ts.URL
	os.Exit(m.Run())
}

func ExampleClient_GetAllWallets() {
	c := itbit.NewClient("key", "secret")

	wallets, err := c.GetAllWallets("userID", 0, 0)
	if err != nil {
		log.Panic(err)
	}

	for _, wallet := range wallets {
		fmt.Printf("%s:\n", wallet.Name)
		for _, balance := range wallet.Balances {
			fmt.Printf("%10.2f %s\t", balance.AvailableBalance, balance.Currency)
		}
		fmt.Printf("\n")
	}
	// Wallet 1:
	// 0.00 USD	      0.00 XBT	      0.00 EUR	      0.00 SGD
	// Wallet 2:
	// 75631.89 USD	 100100.03 XBT	 100000.00 EUR	 100000.00 SGD
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
