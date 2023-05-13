package buffer

import (
	"fmt"
	"testing"
)

var (
	defaultPath = "../base_config/default.yaml"
)

func TestCreateBufferFromConfigFile(t *testing.T) {
	b := CreateBufferFromConfigFile(defaultPath)
	fmt.Println(b)
}

func TestBuffer_DebugPrint(t *testing.T) {
	b := CreateBufferFromConfigFile(defaultPath)
	b.DebugPrint()
}

func TestBuffer_SetService(t *testing.T) {
	b := CreateBufferFromConfigFile(defaultPath)
	b.DebugPrint()

	b.SetService(2, Config{
		Id:          2,
		ServiceName: "NewName",
		Host:        "",
		Port:        "",
		Scheme:      "",
		BufferSize:  0,
	})
	b.DebugPrint()
}

func TestBuffer_GetService(t *testing.T) {
	b := CreateBufferFromConfigFile(defaultPath)
	b.DebugPrint()

	service := b.GetService("ROOT_SERVICE")
	fmt.Println("\n", service)
}
