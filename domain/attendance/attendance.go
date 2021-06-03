package attendance

import (
	"time"
)

type Attendance struct {
	Datetime time.Time
	Status   []*AttendanceStatus
}

func NewAttendanceRecord(status string) (*Attendance, error) {
	attendanceStatus, err := NewAttendanceStatus(status)
	if err != nil {
		return nil, err
	}

	return &Attendance{
		Datetime: time.Now(),
		Status:   []*AttendanceStatus{attendanceStatus},
	}, nil
}
