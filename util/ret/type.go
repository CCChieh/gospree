package ret

import "time"

type baseResult struct {
	Time time.Time `json:"time"`
}

type okResult struct {
	baseResult
	Data interface{} `json:"data"`
}

type errorResult struct {
	baseResult
	Code   int         `json:"code"`
	Errors interface{} `json:"errors"`
}
