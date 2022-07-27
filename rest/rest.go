package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/Goscord/goscord/rest/ratelimit"
)

type Client struct {
	token string
}

func NewClient(token string) *Client {
	return &Client{token: token}
}

func (c *Client) Request(endpoint, method string, data io.Reader, contentType string) ([]byte, error) {
	var req *http.Request

	method = strings.ToUpper(method)
	url := strings.ToLower(BaseUrl + endpoint)
	req, err := http.NewRequest(method, url, data)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("User-Agent", "DiscordBot (https://github.com/Goscord/goscord, 1.0.0)")
	req.Header.Set("Authorization", fmt.Sprintf("Bot %s", c.token))

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var body []byte

	body, err = io.ReadAll(resp.Body)

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

		body, err = c.Request(endpoint, method, data, contentType)
	case 401:
		return nil, errors.New("an invalid token was provided")
	}

	return body, nil
}
