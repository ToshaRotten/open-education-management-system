package configurator_http_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestClient_CheckService(t *testing.T) {
	cli := New()
	if cli.CheckConfigurator() {
		fmt.Println("available")
	} else {
		fmt.Println("not available")
	}
}

func TestClient_GetConfig(t *testing.T) {
	cli := New()
	conf, err := cli.GetConfig("ROOT_SERVICE")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(conf)
}

func TestServer_GetConfig(t *testing.T) {
	cli := http.Client{}
	conf := Config{
		Id:          0,
		ServiceName: "AUTH_SERVICE",
		Host:        "",
		Port:        "",
		Scheme:      "",
		BufferSize:  0,
	}
	data, err := json.Marshal(conf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	req, err := http.NewRequest("GET", "http://localhost:8000/config", bytes.NewReader(data))
	if err != nil {
		fmt.Println(err)
	}
	resp, err := cli.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	err = json.Unmarshal(data, &conf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(conf)
}
