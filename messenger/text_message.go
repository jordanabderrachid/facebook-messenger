package messenger

import (
	"bytes"
	"encoding/json"
	"io"
)

const maxTextLength = 320

type textMessage struct {
	recipientID string
	text        string
}

type facebookTextMessage struct {
	Recipient recipient `json:"recipient"`
	Message   message   `json:"message"`
}

type recipient struct {
	RecipientID string `json:"id"`
}

type message struct {
	Text string `json:"text"`
}

func newTextMessage(recipientID string, text string) textMessage {
	if len(text) > maxTextLength {
		text = text[:maxTextLength]
	}

	return textMessage{
		recipientID: recipientID,
		text:        text,
	}
}

func (tm textMessage) toBody() (io.Reader, error) {
	ftm := facebookTextMessage{
		Recipient: recipient{
			RecipientID: tm.recipientID,
		},
		Message: message{
			Text: tm.text,
		},
	}

	data, err := json.Marshal(ftm)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(data), nil
}
