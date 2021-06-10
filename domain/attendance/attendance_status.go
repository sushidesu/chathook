package attendance

import (
	"errors"
)

type AttendanceStatus struct {
	Value string
}

var AVAILABLE_STATUS = []string{"ENTER", "LEAVE"}

// value は "ENTER" or "LEAVE"
func NewAttendanceStatus(value string) (*AttendanceStatus, error) {
	for _, v := range AVAILABLE_STATUS {
		if v != value {
			return nil, errors.New("invalid status")
		}
	}
	return &AttendanceStatus{Value: value}, nil
}
