package messenger

import "fmt"

// Response comment
type Response struct {
	RecipientID string `json:"recipient_id"`
	MessageID   string `json:"message_id"`
}

func (r Response) String() string {
	return fmt.Sprintf("recipientID=%s, messageID=%s", r.RecipientID, r.MessageID)
}

type badRequestResponse struct {
	Error struct {
		Message   string `json:"message"`
		Type      string `json:"type"`
		Code      int    `json:"code"`
		FBTraceID string `json:"fbtrace_id"`
	} `json:"error"`
}

func (brr badRequestResponse) toError() error {
	return fmt.Errorf("Bad request, message=%s, type=%s, code=%d, fbTraceID=%s",
		brr.Error.Message, brr.Error.Type, brr.Error.Code, brr.Error.FBTraceID)
}
