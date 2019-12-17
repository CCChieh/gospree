package ret

import "time"

type baseResult struct {
	Time time.Time `json:"time"`
	Code int       `json:"code"`
}

type okResult struct {
	baseResult
	Data interface{} `json:"data,omitempty"`
}

type errorResult struct {
	baseResult
	Message interface{} `json:"message,omitempty"`
}
