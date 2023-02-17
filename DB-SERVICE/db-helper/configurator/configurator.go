package configurator

type Config struct {
	Host int
}

//New ..
func New() *Config {
	return &Config{
		Host: 0,
	}
}

func NewFromFile(path string) *Config {
	var c Config
	return &c
}
