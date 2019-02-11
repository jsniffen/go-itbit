package itbit

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTicker(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, client")
	}))
	defer s.Close()

	ItBitEndpoint := s.URL
	fmt.Println(ItBitEndpoint)

	m := newMarketDataService(s.Client())
	resp, _ := m.GetTicker(BitcoinUSDollar)
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	t.Errorf("%s", string(b))
}
