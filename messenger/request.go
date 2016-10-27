package messenger

type requestPayload struct {
	Recipient *recipient `json:"recipient,omitempty"`
	Message   *message   `json:"message,omitempty"`
}

type recipient struct {
	RecipientID string `json:"id,omitempty"`
}

type message struct {
	Text         string       `json:"text,omitempty"`
	Attachment   *attachment  `json:"attachment,omitempty"`
	QuickReplies []quickReply `json:"quick_replies,omitempty"`
}

type quickReply struct {
	ContentType string `json:"content_type,omitempty"`
	Title       string `json:"title,omitempty"`
	Payload     string `json:"payload,omitempty"`
	ImageURL    string `json:"image_url,omitempty"`
}

type attachment struct {
	Type    string   `json:"type,omitempty"`
	Payload *payload `json:"payload,omitempty"`
}

type payload struct {
	URL           string       `json:"url,omitempty"`
	TemplateType  string       `json:"template_type,omitempty"`
	Text          string       `json:"text,omitempty"`
	RecipientName string       `json:"recipient_name,omitempty"`
	OrderNumber   string       `json:"order_number,omitempty"`
	Currency      string       `json:"currency,omitempty"`
	PaymentMethod string       `json:"payment_method,omitempty"`
	OrderURL      string       `json:"order_url,omitempty"`
	Timestamp     string       `json:"timestamp,omitempty"`
	Elements      []element    `json:"elements,omitempty"`
	Address       *address     `json:"address,omitempty"`
	Summary       *summary     `json:"summary,omitempty"`
	Adjustments   []adjustment `json:"adjustments,omitempty"`
}

type element struct {
	Title    string   `json:"title,omitempty"`
	ItemURL  string   `json:"item_url,omitempty"`
	ImageURL string   `json:"image_url,omitempty"`
	Subtitle string   `json:"subtitle,omitempty"`
	Quantity int      `json:"quantity,omitempty"`
	Price    float64  `json:"price,omitempty"`
	Currency string   `json:"currency,omitempty"`
	Buttons  []button `json:"buttons,omitempty"`
}

type button struct {
	Type                string          `json:"type,omitempy"`
	Title               string          `json:"title,omitempty"`
	URL                 string          `json:"url,omitempy"`
	Payload             string          `json:"payload,omitempy"`
	WebviewHeightRatio  string          `json:"webview_height_ratio,omitempty"`
	MessengerExtensions bool            `json:"messenger_extensions,omitempty"`
	FallbackURL         string          `json:"fallback_url"`
	PaymentSummary      *paymentSummary `json:"payment_summary,omitempty"`
}

type address struct {
	Street1    string `json:"street_1,omitempty"`
	Street2    string `json:"street_2,omitempty"`
	City       string `json:"city,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
	State      string `json:"state,omitempty"`
	Country    string `json:"country,omitempty"`
}

type summary struct {
	Subtotal     float64 `json:"subtotal,omitempty"`
	ShippingCost float64 `json:"shipping_cost,omitempty"`
	TotalTax     float64 `json:"total_tax,omitempty"`
	TotalCost    float64 `json:"total_cost,omitempty"`
}

type adjustment struct {
	Name   string  `json:"name,omitempty"`
	Amount float64 `json:"amount,omitempty"`
}

type paymentSummary struct {
	Currency          string   `json:"currency,omitempty"`
	PaymentType       string   `json:"payment_type,omitempty"`
	MerchantName      string   `json:"merchant_name,omitempty"`
	RequestedUserInfo []string `json:"requested_user_info,omitempty"`
	PriceList         []price  `json:"price_list,omitempty"`
}

type price struct {
	Label  string `json:"label,omitempty"`
	Amount string `json:"amount,omitempty"`
}
