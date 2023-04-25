package hello

type Config struct {
	prefix string `mapstructure:"prefix"`
}

func (c *Config) InitDefaults() {
	if c.prefix == "" {
		c.prefix = "Hello, "
	}
}
