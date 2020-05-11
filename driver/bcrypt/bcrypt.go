package bcrypt

import (
	"github.com/jameskeane/bcrypt"
	"log"
)

type Config struct {
	// Number of bcrypt rounds
	Complexity int
}

type Driver struct {
	// Configuration for the bcrypt driver
	Config *Config
	// Stored salt
	salt string
}

func New(config ...Config) *Driver {
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	}
	// Set to default complexity if incorrectly specified
	if cfg.Complexity == 0 {
		cfg.Complexity = bcrypt.DefaultRounds
	}
	salt, err := bcrypt.Salt(cfg.Complexity)
	if err != nil {
		log.Fatalf("Error when generating salt: %v", err)
	}
	return &Driver{
		Config: &cfg,
		salt:   salt,
	}
}

func (d *Driver) CreateHash(password string) (hash string, err error) {
	// hash and verify a password with a static salt
	return bcrypt.Hash(password, d.salt)
}

func (d *Driver) MatchHash(password string, hash string) (match bool, err error) {
	return bcrypt.Match(password, hash), nil
}
