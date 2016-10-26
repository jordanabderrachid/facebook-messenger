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

func newTextMessage(recipientID, text string) textMessage {
	if len(text) > maxTextLength {
		text = text[:maxTextLength]
	}

	return textMessage{
		recipientID: recipientID,
		text:        text,
	}
}

func (tm textMessage) toBody() (io.Reader, error) {
	p := requestPayload{
		Recipient: &recipient{
			RecipientID: tm.recipientID,
		},
		Message: &message{
			Text: tm.text,
		},
	}

	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(data), nil
}
