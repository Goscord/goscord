package rest

import (
	"net/http"
	"strings"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"errors"
	"time"
	"github.com/Seyz123/yalis/rest/ratelimit"
)

type Client struct {
	token string
	http *http.Client
}

func NewClient(token string) *Client {
	return &Client{
		token: token,
		http: &http.Client{},
	}
}

func (c *Client) Request(endpoint, method string, data []byte) ([]byte, error) {
	method = strings.ToUpper(method)
	url := strings.ToLower(BaseUrl + endpoint)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))

	if err != nil {
		panic(err)
	}

	req.Header.Set("User-Agent", "DiscordBot (https://github.com/Seyz123/yalis, 1.0.0)")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bot %s", c.token))

	resp, err := c.http.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	
	var body []byte
	
	body, err := ioutil.ReadAll(resp.Body)

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
		ratelimit, err := ratelimit.NewRateLimit(body)

		if err != nil {
			return nil, err
		}

		// ToDo : Handle ratelimit cleaner lmao

		time.Sleep(ratelimit.RetryAfter)
		
		body, err = c.Request(endpoint, method, data)
	case 401:
		return nil, errors.New("An invalid token was provided")
	}

	return body, nil
}