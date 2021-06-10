package recordattendance

type IAirtable_Client interface {
	CreateRecord(record CreateAttendanceRecord)
}
