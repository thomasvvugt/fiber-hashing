package argon2id

import (
	"github.com/alexedwards/argon2id"
)

type Config struct {
	// Argon2id configuration
	Params *argon2id.Params
}

type Driver struct {
	// Configuration for the argon2id driver
	Config *Config
}

func New(config ...Config) *Driver {
	var cfg Config
	cfg.Params = argon2id.DefaultParams
	if len(config) > 0 {
		cfg = config[0]
	}
	return &Driver{Config: &cfg}
}

func (d *Driver) CreateHash(password string) (hash string, err error) {
	// CreateHash returns a Argon2id hash of a plain-text password using the
	// provided algorithm parameters. The returned hash follows the format used
	// by the Argon2 reference C implementation and looks like this:
	// $argon2id$v=19$m=65536,t=3,p=2$c29tZXNhbHQ$RdescudvJCsgt3ub+b+dWRWJTmaaJObG
	return argon2id.CreateHash(password, d.Config.Params)
}

func (d *Driver) MatchHash(password string, hash string) (match bool, err error) {
	// ComparePasswordAndHash performs a constant-time comparison between a
	// plain-text password and Argon2id hash, using the parameters and salt
	// contained in the hash. It returns true if they match, otherwise it returns
	// false.
	return argon2id.ComparePasswordAndHash(password, hash)
}
