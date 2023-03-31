package schedule

import "main/GroupSchedule"

type Schedule struct {
	GroupSchedules []GroupSchedule.GroupSchedule
}

func New() *Schedule {
	return &Schedule{
		GroupSchedules: nil,
	}
}
