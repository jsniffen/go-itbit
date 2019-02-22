package itbit_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/juliansniff/go-itbit/itbit"
)

func TestMain(m *testing.M) {
	r := mux.NewRouter()
	r.HandleFunc("/wallet_transfers", handleNewWalletTransfer).Methods(http.MethodPost)
	r.HandleFunc("/markets/{tickerSymbol}/ticker", handleGetTicker).Methods(http.MethodGet)
	r.HandleFunc("/markets/{tickerSymbol}/trades", handleGetRecentTrades).Methods(http.MethodGet)
	r.HandleFunc("/markets/{tickerSymbol}/order_book", handleGetOrderBook).Methods(http.MethodGet)
	r.HandleFunc("/wallets/{id}/balances/{currency}", handleGetWalletBalance).Methods(http.MethodGet)
	r.HandleFunc("/wallets/{id}/orders/{orderID}", handleGetOrder).Methods(http.MethodGet)
	r.HandleFunc("/wallets/{id}/orders/{orderID}", handleCancelOrder).Methods(http.MethodDelete)
	r.HandleFunc("/wallets/{id}/orders", handleGetOrders).Methods(http.MethodGet)
	r.HandleFunc("/wallets/{id}/orders", handleCreateNewOrder).Methods(http.MethodPost)
	r.HandleFunc("/wallets/{id}/trades", handleGetFundingHistory).Methods(http.MethodGet)
	r.HandleFunc("/wallets/{id}", handleGetWallet).Methods(http.MethodGet)
	r.HandleFunc("/wallets", handleGetAllWallets).Methods(http.MethodGet)
	r.HandleFunc("/{id}", handleCreateNewWallet).Methods(http.MethodPost)
	ts := httptest.NewServer(r)
	defer ts.Close()

	itbit.Endpoint = ts.URL
	os.Exit(m.Run())
}

func handleGetFundingHistory(w http.ResponseWriter, r *http.Request) {
	response := `
		{
			"totalNumberOfRecords": "2",
			"currentPageNumber": "1",
			"latestExecutionId": "332",
			"recordsPerPage": "50",
			"fundingHistory": [
				{
					"bankName": "fb6",
					"withdrawalId": 94,
					"holdingPeriodCompletionDate": "2015-03-21T17:37:39.9170000",
					"time": "2015-03-18T17:37:39.9170000",
					"currency": "EUR",
					"transactionType": "Withdrawal",
					"amount": "1.00000000",
					"walletName": "Wallet",
					"status": "relayed"
				},
				{
					"destinationAddress": "mfsANnSPCgeRZoc8KYwCA71mmQMuKjUgFJ",
					"txnHash": "b77ded847997fa52cb340aa65239990d71e02ce335430bab19c20a4c3e84e48f",
					"time": "2015-02-04T18:52:39.1270000",
					"currency": "XBT",
					"transactionType": "Deposit",
					"amount": "14.89980000",
					"walletName": "Wallet",
					"status": "completed"
				}
			]
		}`
	fmt.Fprintf(w, response)
}

func handleNewWalletTransfer(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, string(b))
}

func handleGetOrderBook(w http.ResponseWriter, r *http.Request) {
	response := `{
		"asks": [
			[ "219.82", "2.19" ],
			[ "219.83", "6.05" ],
			[ "220.19", "17.59" ],
			[ "220.52", "3.36" ],
			[ "220.53", "33.46" ]
		],
		"bids": [
			[ "219.40", "17.46" ],
			[ "219.13", "53.93" ],
			[ "219.08", "2.20" ],
			[ "218.58", "98.73" ],
			[ "218.20", "3.37" ]
		]
	}`
	fmt.Fprintf(w, response)
}

func handleGetRecentTrades(w http.ResponseWriter, r *http.Request) {
	response := `{
		"count": 3,
		"recentTrades": [
			{
				"timestamp": "2015-05-22T17:45:34.7570000Z",
				"matchNumber": "5CR1JEUBBM8J",
				"price": "351.45000000",
				"amount": "0.00010000"
			},
			{
				"timestamp": "2015-05-22T17:01:08.4270000Z",
				"matchNumber": "5CR1JEUBBM8F",
				"price": "352.00000000",
				"amount": "0.00010000"
			},
			{
				"timestamp": "2015-05-22T17:01:04.8630000Z",
				"matchNumber": "5CR1JEUBBM8C",
				"price": "351.45000000",
				"amount": "0.00010000"
			}
		]
	}`
	fmt.Fprintf(w, response)
}

func handleGetTicker(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tickerSymbol := vars["tickerSymbol"]
	response := fmt.Sprintf(`{
		  "pair": "%s",
		  "bid": "622",
		  "bidAmt": "0.0006",
		  "ask": "641.29",
		  "askAmt": "0.5",
		  "lastPrice": "618.00000000",
		  "lastAmt": "0.00040000",
		  "volume24h": "0.00040000",
		  "volumeToday": "0.00040000",
		  "high24h": "618.00000000",
		  "low24h": "618.00000000",
		  "highToday": "618.00000000",
		  "lowToday": "618.00000000",
		  "openToday": "618.00000000",
		  "vwapToday": "618.00000000",
		  "vwap24h": "618.00000000",
		  "serverTimeUTC": "2014-06-24T20:42:35.6160000Z"
		}`, tickerSymbol)
	fmt.Fprintf(w, response)
}

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

func handleGetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	response := fmt.Sprintf(`
		{
			"id": "5b1f51e1-3f38-4d64-918c-45c5848c76fb",
			"walletId": "%s",
			"side": "buy",
			"instrument": "XBTUSD",
			"type": "limit",
			"currency": "XBT",
			"amount": "2.5",
			"price": "650",
			"amountFilled": "1",
			"volumeWeightedAveragePrice": "650",
			"createdTime": "2014-02-11T17:05:15Z",
			"status": "open",
			"metadata": {},
			"clientOrderIdentifier": null
		}
	`, id)
	fmt.Fprintf(w, response)
}

func handleGetOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	response := fmt.Sprintf(`
		[
			{
				"id": "5b1f51e1-3f38-4d64-918c-45c5848c76fb",
				"walletId": "%s",
				"side": "buy",
				"instrument": "XBTUSD",
				"type": "limit",
				"currency": "XBT",
				"amount": "2.5",
				"price": "650",
				"amountFilled": "1",
				"volumeWeightedAveragePrice": "650",
				"createdTime": "2014-02-11T17:05:15Z",
				"status": "open",
				"metadata": {},
				"clientOrderIdentifier": null
			}
		]
	`, id)
	fmt.Fprintf(w, response)
}

func handleCreateNewOrder(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, string(b))
}

func handleCancelOrder(w http.ResponseWriter, r *http.Request) {
	return
}
