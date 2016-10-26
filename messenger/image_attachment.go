package messenger

import (
	"bytes"
	"encoding/json"
	"io"
)

const imageAttachmentType = "image"

type imageAttachment struct {
	recipientID string
	imageURL    string
}

func newImageAttachment(recipientID, imageURL string) imageAttachment {
	return imageAttachment{
		recipientID: recipientID,
		imageURL:    imageURL,
	}
}

func (ia imageAttachment) toBody() (io.Reader, error) {
	p := requestPayload{
		Recipient: &recipient{
			RecipientID: ia.recipientID,
		},
		Message: &message{
			Attachment: &attachment{
				Type: imageAttachmentType,
				Payload: &payload{
					URL: ia.imageURL,
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
