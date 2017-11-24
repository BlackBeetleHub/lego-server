package json

type ErrorResponse interface {
	GetTypeError() string
}

type ErrorMalformedRequest struct {
	request string
	response string
}

type ErrorNotConnect struct {
	message string
}
