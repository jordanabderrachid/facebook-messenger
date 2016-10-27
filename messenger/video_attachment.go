package messenger

import (
	"bytes"
	"encoding/json"
	"io"
)

type videoAttachment struct {
	recipientID string
	videoURL    string
}

func newVideoAttachment(recipientID, videoURL string) videoAttachment {
	return videoAttachment{
		recipientID: recipientID,
		videoURL:    videoURL,
	}
}

func (va videoAttachment) toBody() (io.Reader, error) {
	p := requestPayload{
		Recipient: &recipient{
			RecipientID: va.recipientID,
		},
		Message: &message{
			Attachment: &attachment{
				Type: videoAttachmentType,
				Payload: &payload{
					URL: va.videoURL,
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
