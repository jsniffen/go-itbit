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

const (
	Bitcoin         = "XBT"
	Etherium        = "ETH"
	Euro            = "EUR"
	SingaporeDollar = "SGD"
	USDollar        = "USD"

	BitcoinUSDollar         = "XBTUSD"
	BitcoinSingaporeDollar  = "XBTSGD"
	BitcoinEuro             = "XBTEUR"
	EtheriumUSDollar        = "ETHUSD"
	EtheriumEuro            = "ETHEUR"
	EtheriumSingaporeDollar = "ETHSGD"
)

var (
	endpoint = "https://api.itbit.com/v1"
	epoch    = func() int64 {
		return time.Now().UnixNano() / 1000000
	}
	nonce = func() int64 {
		return time.Now().UnixNano() / 1000000
	}
)

type Client struct {
	httpClient *http.Client
	key        string
	secret     string
}

func NewClient(key, secret string) *Client {
	return &Client{
		httpClient: &http.Client{},
		key:        key,
		secret:     secret,
	}
}

func (c *Client) SetHttpClient(client *http.Client) {
	c.httpClient = client
}

func (c *Client) SetKey(key string) {
	c.key = key
}

func (c *Client) SetSecret(secret string) {
	c.secret = secret
}

func (c *Client) signRequest(r *http.Request) error {
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
	mac := hmac.New(sha512.New, []byte(c.secret))
	_, err = mac.Write([]byte(r.URL.String() + string(hash[:])))
	if err != nil {
		return fmt.Errorf("error computing HMAC hash: %v", err)
	}
	sum := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	authHeader := fmt.Sprintf("%s:%s", c.key, sum)

	r.Header.Set("Authorization", authHeader)
	r.Header.Set("X-Auth-Timestamp", timestamp)
	r.Header.Set("X-Auth-Nonce", nonce)
	r.Header.Set("Content-Type", "application/json")

	return nil
}
