package trading

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func init() {
	epoch = func() int64 {
		return 0
	}
	nonce = func() int64 {
		return 0
	}
}

func mockRequest(method, URL, body string) *http.Request {
	reader := ioutil.NopCloser(bytes.NewBufferString(body))
	r, _ := http.NewRequest(method, URL, reader)
	return r
}

func TestSignRequest(t *testing.T) {
	s := &TradingService{
		httpClient: &http.Client{},
		endpoint:   "endpoint",
		key:        "key",
		secret:     "secret",
	}

	tests := map[*http.Request](map[string]string){
		mockRequest(http.MethodGet, "get-endpoint", ""): map[string]string{
			"Authorization":    "key:PwcuTHG+Yzkybs1Q7ReKJg86jqO+eWnRsSBYgmuvOby2WQ+TVGTQYUogt5QslBASlmfNQxXMmvfZV60+yK/8NQ==",
			"X-Auth-Timestamp": "0",
			"X-Auth-Nonce":     "0",
			"Content-Type":     "application/json",
		},
		mockRequest(http.MethodGet, "get-endpoint", "get body"): map[string]string{
			"Authorization":    "key:nMu5IJo9LooMbrNG1Xk4wUz8ru5j/6k8yZ7Gzmtbm+RRwWI+l9AkVLp5ksOxUBgmEmAblzg7miX9BcRJ1VuZZw==",
			"X-Auth-Timestamp": "0",
			"X-Auth-Nonce":     "0",
			"Content-Type":     "application/json",
		},
		mockRequest(http.MethodPost, "post-endpoint", ""): map[string]string{
			"Authorization":    "key:tmfXBl4nDuXOI1mn3NgRZcvI5oEtNgnzcKtRAeTZVyv9hn6nOS3hGzARomLMm7gtlFI9FWSGLVr5krt5Zx6V3Q==",
			"X-Auth-Timestamp": "0",
			"X-Auth-Nonce":     "0",
			"Content-Type":     "application/json",
		},
		mockRequest(http.MethodPost, "post-endpoint", "post body"): map[string]string{
			"Authorization":    "key:rkSaNY9lucJWhpPLNBcpXaqS2ee7s8kfwdlYAkh2z/CHK++aMRTj8kdc9BighCAETGCf7bAh2XrZpK5DikMZsg==",
			"X-Auth-Timestamp": "0",
			"X-Auth-Nonce":     "0",
			"Content-Type":     "application/json",
		},
		mockRequest(http.MethodDelete, "delete-endpoint", ""): map[string]string{
			"Authorization":    "key:fTHYTHf9/Kd2kAGe09n6qQ3c/D8VdDQKNHl27TqIaFEC4umGaQvoqJS84WJ7sxQUV6E4FbJhzVisjtv7jYerGA==",
			"X-Auth-Timestamp": "0",
			"X-Auth-Nonce":     "0",
			"Content-Type":     "application/json",
		},
		mockRequest(http.MethodDelete, "delete-endpoint", "delete body"): map[string]string{
			"Authorization":    "key:7VjGDwwLLkDD/ebOYuZf1wo9P19nbYjevEuIcCvO5u5W1B6X01eB+gobKo2V8ADTwlXRSieK6sFhV7lH1PUkhQ==",
			"X-Auth-Timestamp": "0",
			"X-Auth-Nonce":     "0",
			"Content-Type":     "application/json",
		},
	}

	for request, headers := range tests {
		err := s.signRequest(request)
		if err != nil {
			t.Errorf("error signing request: %v", err)
		}

		for key, expected := range headers {
			if got := request.Header.Get(key); got != expected {
				t.Errorf("Expected %s to be: %v, got %v", key, expected, got)
			}
		}
	}
}
