package attendance

import (
	"time"
)

type Attendance struct {
	Datetime  time.Time
	Status    AttendanceStatus
	PlaceType AttendancePlaceType
}

func NewAttendanceRecord(status string, place string) (*Attendance, error) {
	attendanceStatus, err := NewAttendanceStatus(status)
	if err != nil {
		return nil, err
	}

	placeType, err := NewAttendancePlaceType(place)
	if err != nil {
		return nil, err
	}

	return &Attendance{
		Datetime:  time.Now(),
		Status:    *attendanceStatus,
		PlaceType: *placeType,
	}, nil
}

func NewAttendanceRecordSpecifyTime(status string, place string, dateTime time.Time) (*Attendance, error) {
	attendanceStatus, err := NewAttendanceStatus(status)
	if err != nil {
		return nil, err
	}

	placeType, err := NewAttendancePlaceType(place)
	if err != nil {
		return nil, err
	}

	return &Attendance{
		Datetime:  dateTime,
		Status:    *attendanceStatus,
		PlaceType: *placeType,
	}, nil
}
