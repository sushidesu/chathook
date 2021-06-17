package parsemessage

import "strings"

type ParseMessageUsecase struct{}

var BODY_LEAVE_SIGNS = []string{"お疲れ様です(bow)"}

func (parseMessage ParseMessageUsecase) Parse(body string) ParseMessageOutput {
	for _, sign := range BODY_LEAVE_SIGNS {
		if strings.Contains(body, sign) {
			output, _ := NewParseMessageOutput("LEAVE_HOME")
			return *output
		}
	}
	output, _ := NewParseMessageOutput("OTHER")
	return *output
}
