package GroupSchedule

import "main/lesson"

type GroupSchedule struct {
	GroupName string
	Week      []Day
}

type Day struct {
	lessons []lesson.Lesson
}
