package itbit

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

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

	endpoint = ts.URL
	c := NewClient("", "")

	got, err := c.GetOrderBook("tickerSymbol")
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
