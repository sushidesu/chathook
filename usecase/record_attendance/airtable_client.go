package recordattendance

type IAirtable_Client interface {
	CreateRecord(record map[string]interface{})
	DATETIME_FORMAT_STRING() string
}
