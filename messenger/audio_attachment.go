package messenger

import (
	"bytes"
	"encoding/json"
	"io"
)

const audioAttachmentType = "audio"

type audioAttachment struct {
	recipientID string
	audioURL    string
}

func newAudioAttachment(recipientID, audioURL string) audioAttachment {
	return audioAttachment{
		recipientID: recipientID,
		audioURL:    audioURL,
	}
}

func (aa audioAttachment) toBody() (io.Reader, error) {
	p := requestPayload{
		Recipient: &recipient{
			RecipientID: aa.recipientID,
		},
		Message: &message{
			Attachment: &attachment{
				Type: audioAttachmentType,
				Payload: &payload{
					URL: aa.audioURL,
				},
			},
		},
	}

	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}
