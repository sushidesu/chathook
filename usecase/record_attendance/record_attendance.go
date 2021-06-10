package recordattendance

import "github.com/sushidesu/chathook/domain/attendance"

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

type CreateAttendanceRecord map[string]interface{}

// レコードを生成して、文字列に変換したものをAirtableに登録する
func (recordAttendance RecordAttendanceUsecase) Record() {
	attendance, _ := attendance.NewAttendanceRecord("ENTER")

	var statusStrings []string
	for _, status := range attendance.Status {
		statusStrings = append(statusStrings, status.Value)
	}

	record := CreateAttendanceRecord{
		"datetime": attendance.Datetime.Format(recordAttendance.airtableClient.DATETIME_FORMAT_STRING()),
		"status":   statusStrings,
	}

	recordAttendance.airtableClient.CreateRecord(record)
}
