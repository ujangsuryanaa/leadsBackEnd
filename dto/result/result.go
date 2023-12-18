package dto

type SuccessResult struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message,omitempty"`
}

type ErrorResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
