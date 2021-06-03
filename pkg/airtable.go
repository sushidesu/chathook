package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IAirtable interface {
	CreateRecord(recordHasStringFields interface{})
}

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

func (airtable Airtable) CreateRecord(record interface{}) {
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
