package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sushidesu/chathook/infra/airtable"
	parsemessage "github.com/sushidesu/chathook/usecase/parse_message"
	recordattendance "github.com/sushidesu/chathook/usecase/record_attendance"
)

func main() {
	godotenv.Load()

	// エアテーブル記録サンプル
	airtable := airtable.Airtable{
		BaseUrl: os.Getenv("AIRTABLE_BASE_URL"),
		ApiKey:  os.Getenv("AIRTABLE_API_KEY"),
	}
	recordAttendanceUsecase := recordattendance.NewRecordAttendanceUsecase(airtable)
	recordAttendanceUsecase.Record("LEAVE", "OFFICE")

	// メッセージパースサンプル
	parseMessageUsecase := parsemessage.ParseMessageUsecase{}
	result := parseMessageUsecase.Parse("こんにちは、お疲れ様です(bow)")
	fmt.Println(result.Type.Value)
}
