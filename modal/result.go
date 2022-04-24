package modal

type Result struct {
	Data interface{} `json:"data"`
}
type ErrResult struct {
	Message string `json:"message"`
}

type TableResult struct {
	Result
	Total int `json:"total"`
}

func Success(data interface{}) *Result {
	return &Result{
		Data: data,
	}
}

func Err(message string) *ErrResult {
	return &ErrResult{
		Message: message,
	}
}

func TableSucc(data interface{}, total int) *TableResult {
	return &TableResult{
		Result{
			Data: data,
		},
		total,
	}
}
