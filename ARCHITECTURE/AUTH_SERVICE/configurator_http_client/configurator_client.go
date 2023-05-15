package configurator_http_client

import "errors"

type Client struct {
	Host string
	Port string
}

func New() *Client {
	var c Client
	return &c
}

func (c *Client) GetConfig() (Config, error) {
	var err error
	err = errors.New("Not connected to configuration service")
	return Config{}, err
}
