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

// レコードを生成して、文字列に変換したものをAirtableに登録する
func (recordAttendance RecordAttendanceUsecase) Record() {
	attendance, _ := attendance.NewAttendanceRecord("ENTER")

	record := CreateAttendanceRecord{
		Datetime:  attendance.Datetime,
		EventType: attendance.Status.Value,
		PlaceType: // TODO
	}

	recordAttendance.airtableClient.CreateRecord(record)
}
