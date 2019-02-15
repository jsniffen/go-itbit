package itbit

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestGetRecentTrades(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	}))
	defer ts.Close()

	endpoint = ts.URL
	c := NewClient("", "")

	got, err := c.GetRecentTrades("tickerSymbol", "")
	if err != nil {
		t.Errorf("error making request: %v", err)
	}

	expected := RecentTradesResponse{
		Count: 3,
		RecentTrades: []struct {
			Timestamp   time.Time `json:"timestamp"`
			MatchNumber string    `json:"matchNumber"`
			Price       float64   `json:"price,string"`
			Amount      float64   `json:"amount,string"`
		}{
			{
				Timestamp:   time.Date(2015, 5, 22, 17, 45, 34, 757000000, time.UTC),
				MatchNumber: "5CR1JEUBBM8J",
				Price:       351.45000000,
				Amount:      0.00010000,
			},
			{
				Timestamp:   time.Date(2015, 5, 22, 17, 1, 8, 427000000, time.UTC),
				MatchNumber: "5CR1JEUBBM8F",
				Price:       352.00000000,
				Amount:      0.00010000,
			},
			{
				Timestamp:   time.Date(2015, 5, 22, 17, 1, 4, 863000000, time.UTC),
				MatchNumber: "5CR1JEUBBM8C",
				Price:       351.45000000,
				Amount:      0.00010000,
			},
		},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}
