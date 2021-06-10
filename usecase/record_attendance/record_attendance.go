package recordattendance

import (
	"time"

	"github.com/sushidesu/chathook/domain/attendance"
)

/*
EventType は　"ENTER" または　"LEAVE"

PlaceType は "OFFICE" または "HOME"
*/
type CreateAttendanceRecord struct {
	Datetime  time.Time
	EventType string
	PlaceType string
}

type IRecordAttendanceUsecase interface {
	Record()
}

type RecordAttendanceUsecase struct {
	airtableClient IAirtable_Client
}

func NewRecordAttendanceUsecase(airtableClient IAirtable_Client) *RecordAttendanceUsecase {
	return &RecordAttendanceUsecase{
		airtableClient: airtableClient,
	}
}

// 指定した eventType, placeType のレコードを Airtable に記録する
//
// eventType: "ENTER" | "LEAVE"
//
// placeType: "OFFICE" | "HOME"
func (recordAttendance RecordAttendanceUsecase) Record(eventType string, placeType string) {
	attendance, err := attendance.NewAttendanceRecord(eventType, placeType)

	if err != nil {
		panic(err)
	}

	record := CreateAttendanceRecord{
		Datetime:  attendance.Datetime,
		EventType: attendance.Status.Value,
		PlaceType: attendance.PlaceType.Value,
	}

	recordAttendance.airtableClient.CreateRecord(record)
}
