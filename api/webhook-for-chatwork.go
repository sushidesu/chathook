package api

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sushidesu/chathook/infra/chatwork"
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
	// Airtableへ結果を保存
}
