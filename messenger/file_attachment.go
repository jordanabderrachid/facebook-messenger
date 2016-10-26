package messenger

import (
	"bytes"
	"encoding/json"
	"io"
)

const fileAttachmentType = "file"

type fileAttachment struct {
	recipientID string
	fileURL     string
}

func newFileAttachment(recipientID, fileURL string) fileAttachment {
	return fileAttachment{
		recipientID: recipientID,
		fileURL:     fileURL,
	}
}

func (fa fileAttachment) toBody() (io.Reader, error) {
	p := requestPayload{
		Recipient: &recipient{
			RecipientID: fa.recipientID,
		},
		Message: &message{
			Attachment: &attachment{
				Type: fileAttachmentType,
				Payload: &payload{
					URL: fa.fileURL,
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
