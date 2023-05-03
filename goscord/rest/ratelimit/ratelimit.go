package ratelimit

import (
	"github.com/bytedance/sonic"
	"net/http"
	"time"
)

// TODO: Implement global rate-limit

type RateLimit struct {
	Message    string  `json:"message"`
	Global     bool    `json:"global"`
	RetryAfter float64 `json:"retry_after"`
	Bucket     string  `json:"-"`

	CreatedAt time.Time `json:"-"`
}

func NewRateLimit(resp *http.Response, data []byte) (*RateLimit, error) {
	var ratelimit *RateLimit

	err := sonic.Unmarshal(data, &ratelimit)

	if err != nil {
		return nil, err
	}

	ratelimit.CreatedAt = time.Now()
	ratelimit.Bucket = resp.Header.Get("X-RateLimit-Bucket")

	return ratelimit, nil
}

func (rl *RateLimit) Wait() {
	<-time.After(time.Duration(rl.RetryAfter) * time.Second)
}
