package market

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestNewServiceEndpoint(t *testing.T) {
	e := "endpoint"
	s := NewService(&http.Client{}, e)

	got := s.endpoint
	expected := e + "/markets/"

	if got != expected {
		t.Errorf("expected: %s, got: %s", expected, got)
	}
}

func TestGetTicker(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := `{
		  "pair": "XBTUSD",
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
		}`
		fmt.Fprintf(w, response)
	}))
	defer ts.Close()

	s := NewService(&http.Client{}, ts.URL)
	got, _, err := s.GetTicker("tickerSymbol")
	if err != nil {
		t.Errorf("error making request: %v", err)
	}

	expected := TickerInfo{
		Pair:          "XBTUSD",
		Bid:           622,
		BidAmt:        0.0006,
		Ask:           641.29,
		AskAmt:        0.5,
		LastPrice:     618.00000000,
		LastAmt:       0.00040000,
		Volume24H:     0.00040000,
		VolumeToday:   0.00040000,
		High24H:       618.00000000,
		Low24H:        618.00000000,
		HighToday:     618.00000000,
		LowToday:      618.00000000,
		OpenToday:     618.00000000,
		VwapToday:     618.00000000,
		Vwap24H:       618.00000000,
		ServerTimeUTC: time.Date(2014, 6, 24, 20, 42, 35, 616000000, time.UTC),
	}

	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}

func TestGetOrderBook(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	}))
	defer ts.Close()

	s := NewService(&http.Client{}, ts.URL)
	got, _, err := s.GetOrderBook("tickerSymbol")
	if err != nil {
		t.Errorf("error making request: %v", err)
	}

	expected := OrderBook{
		Asks: [][]float64{
			[]float64{219.82, 2.19},
			[]float64{219.83, 6.05},
			[]float64{220.19, 17.59},
			[]float64{220.52, 3.36},
			[]float64{220.53, 33.46},
		},
		Bids: [][]float64{
			[]float64{219.40, 17.46},
			[]float64{219.13, 53.93},
			[]float64{219.08, 2.20},
			[]float64{218.58, 98.73},
			[]float64{218.20, 3.37},
		},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}
