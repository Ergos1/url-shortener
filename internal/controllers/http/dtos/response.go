package dtos

type Response struct {
	Data interface{} `json:"data,omitempty"`
	Err  string      `json:"error,omitempty"`
}
