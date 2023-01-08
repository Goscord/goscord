package ratelimit

import (
	"time"

	"github.com/goccy/go-json"
)

type RateLimit struct {
	Message    string        `json:"message"`
	RetryAfter time.Duration `json:"retry_after"`
	Global     bool          `json:"global"`
}

func NewRateLimit(data []byte) (*RateLimit, error) {
	var ratelimit *RateLimit

	err := json.Unmarshal(data, &ratelimit)

	if err != nil {
		return nil, err
	}

	return ratelimit, nil
}
