package modal

import "net/http"

type Result struct {
	Statue int         `json:"statue"`
	Data   interface{} `json:"data"`
}
type ErrResult struct {
	Statue       int    `json:"statue"`
	ErrorMessage string `json:"errorMessage"`
}

type TableResult struct {
	Result
	Total int `json:"total"`
}

func Success(data interface{}) *Result {
	return &Result{
		Statue: http.StatusOK,
		Data:   data,
	}
}

func Err(statue int, message string) *ErrResult {
	return &ErrResult{
		Statue:       statue,
		ErrorMessage: message,
	}
}

func TableSucc(data interface{}, total int) *TableResult {
	return &TableResult{
		Result{
			Statue: http.StatusOK,
			Data:   data,
		},
		total,
	}
}
