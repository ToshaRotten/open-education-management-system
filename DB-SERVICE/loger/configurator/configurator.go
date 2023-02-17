package configurator

type Config struct {
	LogLevel int
}

//New ..
func New() *Config {
	return &Config{
		LogLevel: 0,
	}
}

func NewFromFile(path string) *Config {
	var c Config
	return &c
}
