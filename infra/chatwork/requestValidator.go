package chatwork

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

type RequestValidator struct {
	WebhookToken string
}

func (validator RequestValidator) Validate(requestBody []byte, signature string) (bool, error) {
	secret, err := base64.StdEncoding.DecodeString(validator.WebhookToken)
	if err != nil {
		return false, err
	}
	mac := hmac.New(sha256.New, secret)
	mac.Write(requestBody)
	digest := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return digest == signature, nil
}
