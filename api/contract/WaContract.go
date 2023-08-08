package contract

type SendWAMessage struct {
	ReceipientType string  `json:"recipient_type"`
	To             string  `json:"to"`
	Type           string  `json:"type"`
	Text           Message `json:"text"`
}

type Message struct {
	Body string `json:"body"`
}
