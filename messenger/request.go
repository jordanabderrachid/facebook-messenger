package messenger

type requestPayload struct {
	Recipient *recipient `json:"recipient,omitempty"`
	Message   *message   `json:"message,omitempty"`
}

type recipient struct {
	RecipientID string `json:"id,omitempty"`
}

type message struct {
	Text       string      `json:"text,omitempty"`
	Attachment *attachment `json:"attachment,omitempty"`
}

type attachment struct {
	Type    string   `json:"type,omitempty"`
	Payload *payload `json:"payload,omitempty"`
}

type payload struct {
	URL string `json:"url,omitempty"`
}
