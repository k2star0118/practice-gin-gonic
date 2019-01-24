package service

import (
	"reflect"
	"sync"
)

type IStore interface {
	Get(path string) (string, error)
	GetKeys(path string) ([]string, error)
	Update(path string, value string) error
	Delete(path string) error
	Close()
}

var (
	containers      = make(map[string]interface{})
	addServiceMutex = &sync.Mutex{}
)

var (
	IStoreType = reflect.TypeOf((*IStore)(nil))
)

func Get(t reflect.Type) interface{} {
	return containers[t.String()]
}

func Add(t reflect.Type, obj interface{}) {
	addServiceMutex.Lock()

	key := t.String()

	if _, ok := containers[key]; !ok {
		containers[key] = obj
	}

	addServiceMutex.Unlock()
}

func Delete(t reflect.Type) {
	addServiceMutex.Lock()

	key := t.String()

	if _, ok := containers[key]; ok {
		delete(containers, key)
	}

	addServiceMutex.Unlock()
}

func Store() IStore {
	return Get(IStoreType).(IStore)
}
