package ret

type baseResult struct {
	Time int64 `json:"time"`
	Code int   `json:"code"`
}

type okResult struct {
	baseResult
	Data interface{} `json:"data,omitempty"`
}

type errorResult struct {
	baseResult
	Message interface{} `json:"message,omitempty"`
}
