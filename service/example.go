package service

import (
	"gin-gonic/model"
	"reflect"
)

var (
	IExampleServiceType = reflect.TypeOf((*IExampleService)(nil))
)

func NewExampleService() IExampleService {
	return &ExampleService{}
}

func GetExampleService() IExampleService {
	return Get(IExampleServiceType).(IExampleService)
}

type IExampleService interface {
	GetExample() (*model.ExampleMessage, error)
	PutExample() (*model.ExampleMessage, error)
}

type ExampleService struct {
}

func (*ExampleService) GetExample() (*model.ExampleMessage, error) {
	exampleMessage := &model.ExampleMessage{
		Message: "Successfully to query get example",
	}
	return exampleMessage, nil
}

func (*ExampleService) PutExample() (*model.ExampleMessage, error) {
	exampleMessage := &model.ExampleMessage{
		Message: "Successfully to query put example",
	}
	return exampleMessage, nil
}
