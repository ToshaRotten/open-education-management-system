package configurator_http_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	Host string
	Port string
	http http.Client
}

func New() *Client {
	var c Client
	//c.Port = "8000"
	//c.Host = "185.21.142.92"
	c.Port = "8000"
	c.Host = "localhost"
	return &c
}

func (c *Client) CheckConfigurator() bool {
	resp, err := c.http.Get("http://" + c.Host + ":" + c.Port + "/health")
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	if string(data) != "" {
		return true
	}
	return false
}

func (c *Client) GetConfig(serviceName string) (Config, error) {
	if c.CheckConfigurator() {
		var config Config
		config.ServiceName = serviceName
		data, err := json.Marshal(config)
		if err != nil {
			fmt.Println(err)
		}
		req, err := http.NewRequest("GET", "http://localhost:8000/config", bytes.NewReader(data))
		if err != nil {
			fmt.Println(err)
			return config, errors.New("not connected to Configurator")
		}
		resp, err := c.http.Do(req)
		if err != nil {
			fmt.Println(err)
		}

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal(data, &config)
		if err != nil {
			fmt.Println(err)
		}
		return config, nil
	}
	return Config{}, errors.New("not connected to Configurator")
}
