package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Goscord/goscord/rest/ratelimit"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	token string
}

func NewClient(token string) *Client {
	return &Client{token: token}
}

func (c *Client) Request(endpoint, method string, data []byte) ([]byte, error) {
	var req *http.Request

	method = strings.ToUpper(method)
	url := strings.ToLower(BaseUrl + endpoint)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "DiscordBot (https://github.com/Goscord/goscord, 1.0.0)")
	req.Header.Set("Authorization", fmt.Sprintf("Bot %s", c.token))

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var body []byte

	body, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var resData map[string]interface{}
	err = json.Unmarshal(body, &resData)

	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case 429:
		rateLimit, err := ratelimit.NewRateLimit(body)

		if err != nil {
			return nil, err
		}

		// ToDo : Handle rateLimit cleaner lmao

		time.Sleep(rateLimit.RetryAfter)

		body, err = c.Request(endpoint, method, data)
	case 401:
		return nil, errors.New("An invalid token was provided")
	}

	return body, nil
}
