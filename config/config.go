package config

import (
	"sync"
)

var (
	config *AppConfig
	once   sync.Once
)

type AppConfig struct {
	HttpEndpoint string
	Dev          bool
	MaxConnCount int
}

func Get() *AppConfig {
	once.Do(func() {
		config = &AppConfig{
			HttpEndpoint: ":8888",
			Dev: false,
			MaxConnCount: 20,
		}
	})

	return config
}
