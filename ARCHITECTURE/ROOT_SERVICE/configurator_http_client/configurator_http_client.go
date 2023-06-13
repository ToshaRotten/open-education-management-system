package configurator_http_client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	host string
	port string
	http http.Client
}

func New() *Client {
	var c Client
	return &c
}

func (c *Client) GetConfig() Config {
	var config Config
	resp, err := c.http.Get("http://" + c.host + c.port + "/config/root-service")
	if err != nil {
		fmt.Println(err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println(err)
	}
	return config
}
