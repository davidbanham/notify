package types

// Email an Email
type Email struct {
	Subject string `json:"subject"`
	Body    Body   `json:"body"`
	From    string `json:"from"`
	To      string `json:"to"`
	ToName  string `json:"toName"`
}

// SMS an SMS
type SMS struct {
	To   string `json:"to"`
	Body string `json:"body"`
	From string `json:"from"`
}

// Body the body of an email
type Body struct {
	Text string `json:"text"`
	HTML string `json:"html"`
}
