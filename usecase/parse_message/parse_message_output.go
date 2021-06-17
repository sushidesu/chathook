package parsemessage

import "errors"

type ParseMessageOutput struct {
	Type ParseMessageType
}

/*
	messageTypeString: "LEAVE_HOME" | "OTHER"
*/
func NewParseMessageOutput(messageTypeString string) (*ParseMessageOutput, error) {
	messageType, err := NewParseMessageType(messageTypeString)
	if err != nil {
		return nil, err
	}
	return &ParseMessageOutput{Type: *messageType}, nil
}

type ParseMessageType struct {
	Value string
}

var AVAILABLE_MESSAGE_TYPES = []string{"LEAVE_HOME", "OTHER"}

func NewParseMessageType(messageType string) (*ParseMessageType, error) {
	for _, typeString := range AVAILABLE_MESSAGE_TYPES {
		if typeString == messageType {
			return &ParseMessageType{messageType}, nil
		}
	}
	return nil, errors.New("invalid message type")
}
