package dto

type HandlerResponse struct {
	StatusCode int
	Message    string
	Data       interface{}
	Error      error
}
