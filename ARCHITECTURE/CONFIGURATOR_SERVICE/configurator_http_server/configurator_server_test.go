package configurator_http_server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/CONFIGURATOR_SERVICE/configurator/models"
	"io"
	"net/http"
	"testing"
)

func TestSearchAvailablePort(t *testing.T) {
	port := SearchAvailablePort()
	fmt.Println(port)
}

func TestSearchIP(t *testing.T) {
	SearchIP()
}

func TestServer_Start(t *testing.T) {
	serv := New()
	serv.Start()
}

func TestServer_GetConfig(t *testing.T) {
	cli := http.Client{}
	conf := models.Config{
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
