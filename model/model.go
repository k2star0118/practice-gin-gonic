package model

type HttpResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ExampleMessage struct {
	Message	string	`json:"message"`
}

type CommonFunctions interface {
	Validate() error
}
