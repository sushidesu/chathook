package api

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sushidesu/chathook/domain/constants"
	"github.com/sushidesu/chathook/infra/airtable"
	"github.com/sushidesu/chathook/infra/chatwork"
	parsemessage "github.com/sushidesu/chathook/usecase/parse_message"
	recordattendance "github.com/sushidesu/chathook/usecase/record_attendance"
)

func WebhookForChatwork(w http.ResponseWriter, r *http.Request) {
	godotenv.Load()

	// ヘッダ、ボディを取得
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	signature := r.Header.Get("X-ChatWorkWebhookSignature")

	// 署名の検証
	validator := chatwork.RequestValidator{WebhookToken: os.Getenv("CHATWORK_WEBHOOK_TOKEN")}
	valid, err := validator.Validate(body, signature)
	if !valid || err != nil {
		return
	}

	// ボディをパース
	parseMessageUsecase := parsemessage.ParseMessageUsecase{}
	parseResult := parseMessageUsecase.Parse(string(body))

	// Airtableへ結果を保存
	// その他のメッセージ
	if parseResult.Type.Value == "OTHER" {
		return
	}

	// 保存準備
	airtableClient := airtable.Airtable{BaseUrl: os.Getenv("AIRTABLE_BASE_URL"), ApiKey: os.Getenv("AIRTABLE_API_KEY")}
	recordAttendanceUsecase := recordattendance.NewRecordAttendanceUsecase(airtableClient)

	// LEAVEメッセージ
	if parseResult.Type.Value == "LEAVE_HOME" {
		recordAttendanceUsecase.RecordSpecifyTime("ENTER", "HOME", constants.RegularTimeEnter())
		recordAttendanceUsecase.Record("LEAVE", "HOME")
		return
	}
}
