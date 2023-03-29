package schedule

import "main/group_schedule"

type Schedule struct {
	GroupSchedules []group_schedule.GroupSchedule
}

func New() *Schedule {
	return &Schedule{
		GroupSchedules: nil,
	}
}
