package attendance

import (
	"errors"
)

type AttendancePlaceType struct {
	Value string
}

var AVAILABLE_PLACE = []string{"OFFICE", "HOME"}

// vaue は　"OFFICE" or "HOME"
func NewAttendancePlaceType(value string) (*AttendancePlaceType, error) {
	for _, v := range AVAILABLE_PLACE {
		if v == value {
			return &AttendancePlaceType{Value: value}, nil
		}
	}
	return nil, errors.New("invalid placeType")
}
