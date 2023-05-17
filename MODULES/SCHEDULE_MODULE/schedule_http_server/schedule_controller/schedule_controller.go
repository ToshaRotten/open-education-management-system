package schedule_controller

import (
	"fmt"
	"io"
	"os"
)

type Controller struct {
}

func New() *Controller {
	var c Controller
	return &c
}

func (c *Controller) Create(data []byte) {
	file, err := os.Create("schedules/temp.dat")
	if err != nil {
		fmt.Println("Create file error:", err)
	}
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Write file error:", err)
	}
}

func (c *Controller) Read() []byte {
	file, err := os.Open("schedules/temp.dat")
	if err != nil {
		fmt.Println("Open file error: ", err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Read file error: ", err)
	}
	return data
}
