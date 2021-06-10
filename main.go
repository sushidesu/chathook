package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sushidesu/chathook/infra/airtable"
	recordattendance "github.com/sushidesu/chathook/usecase/record_attendance"
)

func main() {
	godotenv.Load()

	airtable := airtable.Airtable{
		BaseUrl: os.Getenv("AIRTABLE_BASE_URL"),
		ApiKey:  os.Getenv("AIRTABLE_API_KEY"),
	}

	recordAttendanceUsecase := recordattendance.NewRecordAttendanceUsecase(airtable)
	recordAttendanceUsecase.Record("LEAVE", "OFFICE")
}
