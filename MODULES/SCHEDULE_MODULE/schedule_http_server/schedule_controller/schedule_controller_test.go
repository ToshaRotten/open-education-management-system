package schedule_controller

import (
	"encoding/json"
	"fmt"
	"main/schedule_http_server/models"
	"testing"
)

var (
	les1 = models.Lesson{
		Id:              0,
		Title:           "",
		Start:           "start",
		End:             "",
		BackgroundColor: "",
		ExtendedProps:   models.ExtendedProps{},
	}
	les2 = models.Lesson{
		Id:              1,
		Title:           "title",
		Start:           "start",
		End:             "",
		BackgroundColor: "",
		ExtendedProps:   models.ExtendedProps{},
	}
)

func TestController_Create(t *testing.T) {
	cont := New()
	var lessons []models.Lesson
	lessons = append(lessons, les1)
	lessons = append(lessons, les2)
	data, err := json.Marshal(&lessons)
	if err != nil {
		fmt.Println("Marshall error: ", err)
	}
	cont.Create(data)
}

func TestController_Read(t *testing.T) {
	cont := New()
	var lessons []models.Lesson
	data := cont.Read()
	err := json.Unmarshal(data, &lessons)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(lessons)
}
