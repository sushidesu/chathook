package attendance

import (
	"errors"
)

type AttendanceStatus struct {
	Value string
}

const enterId = "recUZn6wg6HPuIwQ9"
const leaveId = "recirLaEnVTE9aGb3"

var idMap = map[string]string{
	"ENTER": enterId,
	"LEAVE": leaveId,
}

func NewAttendanceStatus(value string) (*AttendanceStatus, error) {
	if value != "ENTER" && value != "LEAVE" {
		return nil, errors.New("invalid status")
	} else {
		return &AttendanceStatus{idMap[value]}, nil
	}
}
