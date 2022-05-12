package livestreaming

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const (
	API_GATEWAY_CHINA_MAINLAND = "https://vcloud.163.com"
	API_GATEWAY_GLOBAL         = "https://api-sea.yunxinvcloud.com"
)

type Client struct {
	appKey    string
	appSecret string

	gatewayURL string
	hc         http.Client
}

type Response[R any] struct {
	Code      int    `json:"code"`
	RequestID string `json:"requestId,omitempty"`
	Msg       string `json:"msg,omitempty"`
	Ret       R      `json:"ret"`
}

func init() {
	rand.Seed(time.Now().UnixMicro())
}
func NewClient(gatewayURL, appKey, appSecret string) *Client {
	c := &Client{
		gatewayURL: gatewayURL,
		appKey:     appKey,
		appSecret:  appSecret,
		hc:         http.Client{},
	}
	return c
}

func (c *Client) doRequest(url string, req any, res any) error {
	b, err := json.Marshal(req)
	if err != nil {
		return err
	}
	r := bytes.NewReader(b)
	request, err := http.NewRequest(http.MethodPost, c.gatewayURL+url, r)
	if err != nil {
		return err
	}
	nonce := fmt.Sprint(rand.Int())
	curTime := fmt.Sprint(time.Now().Unix())
	checksumStr := c.appSecret + nonce + curTime
	sha1Bytes := sha1.Sum([]byte(checksumStr))
	checksum := hex.EncodeToString(sha1Bytes[:])
	// fmt.Printf("checksum: %s, %s", checksumStr, checksum)

	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("AppKey", c.appKey)
	request.Header.Set("Nonce", nonce)
	request.Header.Set("CurTime", curTime)
	request.Header.Set("CheckSum", checksum)
	response, err := c.hc.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(res)
}
