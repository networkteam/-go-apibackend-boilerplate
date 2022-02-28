package mail

import "myvendor.mytld/myproject/backend/domain"

// Config stores mail specific configuration
type Config struct {
	// Embed base config for easier passing around
	domain.Config

	DefaultFrom string
}

func DefaultConfig(c domain.Config) Config {
	config := Config{
		Config: c,
	}

	config.DefaultFrom = "app@myproject.mytld"

	return config
}
