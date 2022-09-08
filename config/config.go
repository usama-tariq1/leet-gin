package config

import "os"

type Env struct {
}

func (config Env) Set(key string, value string) {
	os.Setenv(key, value)
}

func (config Env) Get(key string) string {
	return os.Getenv(key)
}
