package ratelimit

import (
	"time"
	"encoding/json"
)

type RateLimit struct {
	Message string `json:"message,omitempty"`
	RetryAfter time.Duration `json:"retry_after,omitempty"`
	Global bool `json:"global,omitempty"`
}

func NewRateLimit(data []byte) (*RateLimit, error) {
	var ratelimit RateLimit

	err := json.Unmarshal(data, &ratelimit)

	if err != nil {
		return nil, err
	}

	return &ratelimit, nil
}