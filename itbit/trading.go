package itbit

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type TradingService struct {
	httpClient *http.Client
	key        string
	secret     string
	endpoint   string
}

var (
	epoch = func() int64 {
		return time.Now().UnixNano() / 1000000
	}
	nonce = func() int64 {
		return time.Now().UnixNano() / 1000000
	}
)

func newTradingService(c *http.Client, key, secret string) *TradingService {
	return &TradingService{
		httpClient: c,
		key:        key,
		secret:     secret,
		endpoint:   endpoint + "/wallets",
	}
}

func (s *TradingService) GetAllWallets(userID string, page, perPage int) ([]Wallet, error) {
	var wallets []Wallet

	if userID == "" {
		return wallets, fmt.Errorf("userID is required, got empty string")
	}

	URL := fmt.Sprintf("%s?userId=%s", s.endpoint, userID)

	if page != 0 {
		URL = fmt.Sprintf("%s?page=%d", URL, page)
	}

	if perPage != 0 {
		URL = fmt.Sprintf("%s?perPage=%d", URL, perPage)
	}

	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return wallets, err
	}

	err = s.signRequest(req)
	if err != nil {
		return wallets, fmt.Errorf("error signing request: %v", err)
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return wallets, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return wallets, fmt.Errorf("error reading response: %v", err)
	}

	if resp.StatusCode != 200 {
		return wallets, fmt.Errorf("%s", string(b))
	}

	err = json.Unmarshal(b, &wallets)
	if err != nil {
		return wallets, fmt.Errorf("error Unmarshalling response into object: %v", err)
	}

	return wallets, nil
}

func (s *TradingService) signRequest(r *http.Request) error {
	timestamp := strconv.FormatInt(epoch(), 10)
	nonce := strconv.FormatInt(nonce(), 10)

	var body []byte
	var err error
	if r.Body == nil {
		body = []byte("")
	} else {
		body, err = ioutil.ReadAll(r.Body)
		if err != nil {
			return fmt.Errorf("error reading request body: %v", err)
		}
	}

	message, err := json.Marshal([]string{r.Method, r.URL.String(), string(body), nonce, timestamp})
	if err != nil {
		return fmt.Errorf("error marshalling authentication message json: %v", err)
	}

	hash := sha256.Sum256([]byte(nonce + string(message)))
	mac := hmac.New(sha512.New, []byte(s.secret))
	_, err = mac.Write([]byte(r.URL.String() + string(hash[:])))
	if err != nil {
		return fmt.Errorf("error computing HMAC hash: %v", err)
	}
	sum := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	authHeader := fmt.Sprintf("%s:%s", s.key, sum)

	r.Header.Set("Authorization", authHeader)
	r.Header.Set("X-Auth-Timestamp", timestamp)
	r.Header.Set("X-Auth-Nonce", nonce)
	r.Header.Set("Content-Type", "application/json")

	return nil
}
