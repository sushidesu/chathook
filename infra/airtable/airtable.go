package airtable

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	recordattendance "github.com/sushidesu/chathook/usecase/record_attendance"
)

// レコードのID等
const AIRTABLE_DATETIME_FORMAT string = "2006-01-02T15:04:05+09:00"

const ENTER_ID = "recUZn6wg6HPuIwQ9"
const LEAVE_ID = "recirLaEnVTE9aGb3"

var ENTER_OR_LEAVE = map[string]string{
	"ENTER": ENTER_ID,
	"LEAVE": LEAVE_ID,
}

const OFFICE_ID = "rec7TEBnzU6iL24fM"
const HOME_ID = "rec6wMQ9d7qeR6HIF"

var OFFICE_OR_HOME = map[string]string{
	"OFFICE": OFFICE_ID,
	"HOME":   HOME_ID,
}

// ------

type IAirtable = recordattendance.IAirtable_Client

type Airtable struct {
	BaseUrl string
	ApiKey  string
}

type RecordMap map[string]interface{}

type FieldObject struct {
	Field RecordMap `json:"fields"`
}

type CreateRecordObject struct {
	Records []FieldObject `json:"records"`
}

func (airtable Airtable) CreateRecord(record recordattendance.CreateAttendanceRecord) {
	fmt.Println("start request...")
	records := []FieldObject{{
		Field: airtable.convert(record),
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

func (airtable Airtable) convert(record recordattendance.CreateAttendanceRecord) RecordMap {
	eventTypeIds := [1]string{ENTER_OR_LEAVE[record.EventType]}
	placeTypeIds := [1]string{OFFICE_OR_HOME[record.PlaceType]}

	return RecordMap{
		"datetime":    record.Datetime.Format(AIRTABLE_DATETIME_FORMAT),
		"eventTypeId": eventTypeIds,
		"placeTypeId": placeTypeIds,
	}
}
