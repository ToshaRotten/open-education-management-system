package configurator_http_server

import (
	"fmt"
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
