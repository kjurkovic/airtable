package wrappers

type Message struct {
	To      string `json:"to"`
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Text    string `json:"text"`
}
