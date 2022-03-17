package modal

type Result struct {
	Statue       int
	Data         interface{}
	ErrorMessage string
}

func Success(data interface{}) *Result {
	return &Result{
		Statue: 200,
		Data:   data,
	}
}

func Err(statue int, message string) *Result {
	return &Result{
		Statue:       statue,
		ErrorMessage: message,
	}
}
