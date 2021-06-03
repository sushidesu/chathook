package airtable

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	recordattendance "github.com/sushidesu/chathook/usecase/record_attendance"
)

const AIRTABLE_DATETIME_FORMAT string = "2006-01-02T15:04:05+00:00"

type IAirtable = recordattendance.IAirtable_Client

type Airtable struct {
	BaseUrl string
	ApiKey  string
}

type FieldObject struct {
	Field interface{} `json:"fields"`
}

type CreateRecordObject struct {
	Records []FieldObject `json:"records"`
}

func (airtable Airtable) CreateRecord(record map[string]interface{}) {
	fmt.Println("start request...")
	records := []FieldObject{{
		Field: record,
	}}
	recordObject := CreateRecordObject{
		Records: records,
	}

	recordJson, _ := json.Marshal(recordObject)
	req, _ := http.NewRequest("POST", airtable.BaseUrl, bytes.NewBuffer((recordJson)))
	req.Header.Set("Authorization", "Bearer "+airtable.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body: ", string(body))
}

func (airtable Airtable) DATETIME_FORMAT_STRING() string {
	return AIRTABLE_DATETIME_FORMAT
}
