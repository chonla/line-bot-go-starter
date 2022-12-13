package line

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"linehook/models"

	"github.com/go-resty/resty/v2"
)

type Line struct {
	channelAccessToken string
	channelSecret      string
	restClient         *resty.Client
}

func New(channelSecret, channelAccessToken string) *Line {
	return &Line{
		channelSecret:      channelSecret,
		channelAccessToken: channelAccessToken,
		restClient:         resty.New(),
	}
}

func (l *Line) VerifyMessage(message, signature []byte) bool {
	hm := hmac.New(sha256.New, []byte(l.channelSecret))
	hm.Write(message)
	expectedHMAC := []byte(base64.StdEncoding.EncodeToString(hm.Sum(nil)))
	return hmac.Equal(signature, expectedHMAC)
}

func (l *Line) Reply(token, message string) error {
	messageReq := &models.ReplyMessage{
		ReplyToken: token,
		Messages: []models.Message{
			{
				Type: "text",
				Text: message,
			},
		},
	}

	_, err := l.restClient.
		R().
		SetBody(messageReq).
		SetHeaders(map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", l.channelAccessToken),
			"Content-Type":  "application/json",
		}).
		Post("https://api.line.me/v2/bot/message/reply")

	if err != nil {
		return err
	}

	return nil
}
