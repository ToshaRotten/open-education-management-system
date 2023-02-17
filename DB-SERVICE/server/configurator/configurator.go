package configurator

type Config struct {
	Host string
	Port string
}

//New ..
func New() *Config {
	return &Config{
		Host: "",
		Port: "8080",
	}
}

func (c *Config) ParseFile(path string) *Config {
	return nil
}
