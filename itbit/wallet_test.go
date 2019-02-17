package itbit_test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/juliansniff/go-itbit/itbit"
)

func TestMain(m *testing.M) {
	r := mux.NewRouter()
	r.HandleFunc("/wallets/{id}/balances/{currency}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		currency := vars["currency"]

		response := fmt.Sprintf(`{
				"currency": "%s",
				"availableBalance": "0",
				"totalBalance": "0"
			}`, currency)
		fmt.Fprintf(w, response)
	})
	ts := httptest.NewServer(r)
	defer ts.Close()

	itbit.Endpoint = ts.URL
	os.Exit(m.Run())
}

// func TestGetAllWallets(t *testing.T) {
// 	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		response := `[
// 			{
// 				"id": "f46eb6b0-2a7f-4c07-898e-856d58568fde",
// 				"userId": "e9c856db-c726-46b2-ace0-a806671105fb",
// 				"name": "user-wallet-01",
// 				"balances": [
// 					{
// 						"currency": "USD",
// 						"availableBalance": "0.00000000",
// 						"totalBalance": "0.00000000"
// 					},
// 					{
// 						"currency": "XBT",
// 						"availableBalance": "0.00000000",
// 						"totalBalance": "0.00000000"
// 					},
// 					{
// 						"currency": "EUR",
// 						"availableBalance": "0.00000000",
// 						"totalBalance": "0.00000000"
// 					},
// 					{
// 						"currency": "SGD",
// 						"availableBalance": "0.00000000",
// 						"totalBalance": "0.00000000"
// 					}
// 				]
// 			},
// 			{
// 				"id": "2224db00-da07-46f7-ba37-b877e1cd455a",
// 				"userId": "e9c856db-c726-46b2-ace0-a806671105fb",
// 				"name": "user-wallet-02",
// 				"balances": [
// 					{
// 						"currency": "USD",
// 						"availableBalance": "75631.88785992",
// 						"totalBalance": "75631.88785999"
// 					},
// 					{
// 						"currency": "XBT",
// 						"availableBalance": "100100.03000000",
// 						"totalBalance": "100100.03000000"
// 					},
// 					{
// 						"currency": "EUR",
// 						"availableBalance": "100000.00000000",
// 						"totalBalance": "100000.00000000"
// 					},
// 					{
// 						"currency": "SGD",
// 						"availableBalance": "100000.00000000",
// 						"totalBalance": "100000.00000000"
// 					}
// 				]
// 			}
// 		]`
// 		fmt.Fprintf(w, response)
// 	}))
// 	defer ts.Close()
//
// 	endpoint = ts.URL
// 	c := NewClient("key", "secret")
//
// 	got, err := c.GetAllWallets("userID", 0, 0)
// 	if err != nil {
// 		t.Errorf("error getting all wallets: %v", err)
// 		return
// 	}
//
// 	expected := []Wallet{
// 		Wallet{
// 			ID:     "f46eb6b0-2a7f-4c07-898e-856d58568fde",
// 			UserID: "e9c856db-c726-46b2-ace0-a806671105fb",
// 			Name:   "user-wallet-01",
// 			Balances: []Balance{
// 				{
// 					Currency:         "USD",
// 					AvailableBalance: 0.00000000,
// 					TotalBalance:     0.00000000,
// 				},
// 				{
// 					Currency:         "XBT",
// 					AvailableBalance: 0.00000000,
// 					TotalBalance:     0.00000000,
// 				},
// 				{
// 					Currency:         "EUR",
// 					AvailableBalance: 0.00000000,
// 					TotalBalance:     0.00000000,
// 				},
// 				{
// 					Currency:         "SGD",
// 					AvailableBalance: 0.00000000,
// 					TotalBalance:     0.00000000,
// 				},
// 			},
// 		},
// 		Wallet{
// 			ID:     "2224db00-da07-46f7-ba37-b877e1cd455a",
// 			UserID: "e9c856db-c726-46b2-ace0-a806671105fb",
// 			Name:   "user-wallet-02",
// 			Balances: []Balance{
// 				{
// 					Currency:         "USD",
// 					AvailableBalance: 75631.88785992,
// 					TotalBalance:     75631.88785999,
// 				},
// 				{
// 					Currency:         "XBT",
// 					AvailableBalance: 100100.03000000,
// 					TotalBalance:     100100.03000000,
// 				},
// 				{
// 					Currency:         "EUR",
// 					AvailableBalance: 100000.00000000,
// 					TotalBalance:     100000.00000000,
// 				},
// 				{
// 					Currency:         "SGD",
// 					AvailableBalance: 100000.00000000,
// 					TotalBalance:     100000.00000000,
// 				},
// 			},
// 		},
// 	}
//
// 	if !reflect.DeepEqual(got, expected) {
// 		t.Errorf("got: %v, expected: %v", got, expected)
// 	}
// }
//
// func TestCreateNewWallet(t *testing.T) {
// 	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		b, _ := ioutil.ReadAll(r.Body)
// 		r.Body.Close()
// 		var body map[string]string
// 		_ = json.Unmarshal(b, &body)
//
// 		name := body["name"]
// 		userID := body["userId"]
//
// 		response := fmt.Sprintf(`{
// 				"id": "walletID",
// 				"userId": "%s",
// 				"name": "%s",
// 				"balances": [
// 					{
// 						"currency": "USD",
// 						"availableBalance": "0.00000000",
// 						"totalBalance": "0.00000000"
// 					},
// 					{
// 						"currency": "XBT",
// 						"availableBalance": "0.00000000",
// 						"totalBalance": "0.00000000"
// 					},
// 					{
// 						"currency": "EUR",
// 						"availableBalance": "0.00000000",
// 						"totalBalance": "0.00000000"
// 					},
// 					{
// 						"currency": "SGD",
// 						"availableBalance": "0.00000000",
// 						"totalBalance": "0.00000000"
// 					}
// 				]
// 			}`, userID, name)
// 		fmt.Fprintf(w, response)
// 	}))
// 	defer ts.Close()
//
// 	endpoint = ts.URL
// 	c := NewClient("key", "secret")
// 	got, err := c.CreateNewWallet("userID", "walletName")
// 	if err != nil {
// 		t.Errorf("%v", err)
// 		return
// 	}
//
// 	expected := Wallet{
// 		ID:     "walletID",
// 		UserID: "userID",
// 		Name:   "walletName",
// 		Balances: []Balance{
// 			{
// 				Currency:         "USD",
// 				AvailableBalance: 0.00000000,
// 				TotalBalance:     0.00000000,
// 			},
// 			{
// 				Currency:         "XBT",
// 				AvailableBalance: 0.00000000,
// 				TotalBalance:     0.00000000,
// 			},
// 			{
// 				Currency:         "EUR",
// 				AvailableBalance: 0.00000000,
// 				TotalBalance:     0.00000000,
// 			},
// 			{
// 				Currency:         "SGD",
// 				AvailableBalance: 0.00000000,
// 				TotalBalance:     0.00000000,
// 			},
// 		},
// 	}
//
// 	if !reflect.DeepEqual(got, expected) {
// 		t.Errorf("got: %v, expected: %v", got, expected)
// 	}
// }
//
// func TestGetWallet(t *testing.T) {
// 	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		path := r.URL.Path
// 		pathArray := strings.Split(path, "/")
// 		walletID := pathArray[len(pathArray)-1]
// 		response := fmt.Sprintf(`{
// 			"id": "%s",
// 			"userId": "userID",
// 			"name": "walletName",
// 			"balances": [
// 				{
// 					"currency": "USD",
// 					"availableBalance": "50000.0000000",
// 					"totalBalance": "50000.0000000"
// 				},
// 				{
// 					"currency": "XBT",
// 					"availableBalance": "100.00000000",
// 					"totalBalance": "100.00000000"
// 				},
// 				{
// 					"currency": "EUR",
// 					"availableBalance": "100000.00000000",
// 					"totalBalance": "100000.00000000"
// 				},
// 				{
// 					"currency": "SGD",
// 					"availableBalance": "515440.88288502",
// 					"totalBalance": "515432.74228603"
// 				}
// 			]
// 		}`, walletID)
// 		fmt.Fprintf(w, response)
// 	}))
// 	defer ts.Close()
//
// 	endpoint = ts.URL
// 	c := NewClient("key", "secret")
//
// 	got, err := c.GetWallet("walletID")
// 	if err != nil {
// 		t.Errorf("error getting wallets: %v", err)
// 		return
// 	}
//
// 	expected := Wallet{
// 		ID:     "walletID",
// 		UserID: "userID",
// 		Name:   "walletName",
// 		Balances: []Balance{
// 			{
// 				Currency:         "USD",
// 				AvailableBalance: 50000.0000000,
// 				TotalBalance:     50000.0000000,
// 			},
// 			{
// 				Currency:         "XBT",
// 				AvailableBalance: 100.00000000,
// 				TotalBalance:     100.00000000,
// 			},
// 			{
// 				Currency:         "EUR",
// 				AvailableBalance: 100000.00000000,
// 				TotalBalance:     100000.00000000,
// 			},
// 			{
// 				Currency:         "SGD",
// 				AvailableBalance: 515440.88288502,
// 				TotalBalance:     515432.74228603,
// 			},
// 		},
// 	}
//
// 	if !reflect.DeepEqual(got, expected) {
// 		t.Errorf("got: %v, expected: %v", got, expected)
// 	}
// }

func ExampleClient_GetWalletBalance() {
	c := itbit.NewClient("key", "secret")
	balance, err := c.GetWalletBalance("walletID", itbit.USDollar)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("wallet has %.2f %s", balance.TotalBalance, balance.Currency)
	// Output: wallet has 0.00 USD
}
