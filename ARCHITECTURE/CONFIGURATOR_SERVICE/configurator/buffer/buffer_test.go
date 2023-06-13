package buffer

import (
	"fmt"
	"github.com/ToshaRotten/open-education-management-system/ARCHITECTURE/CONFIGURATOR_SERVICE/configurator/models"
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

	b.SetService(2, models.Config{
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
