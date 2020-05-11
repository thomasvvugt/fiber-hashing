package hashing

import (
	"github.com/thomasvvugt/fiber-hashing/driver/argon2id"
)

// Driver interface implemented by drivers
type Driver interface {
	CreateHash(password string) (hash string, err error)
	MatchHash(password string, hash string) (match bool, err error)
}

type Config struct {
	Driver Driver
}

// New ...
func New(config ...Config) Driver {
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	}
	if cfg.Driver == nil {
		cfg.Driver = argon2id.New()
	}
	return cfg.Driver
}
