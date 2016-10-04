package handlers

type JsonError struct {
	Message string `json:"message"`
}

func NewJsonError(m string) *JsonError {
	return &JsonError{Message: m}
}
