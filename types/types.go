package types

// Email an Email
type Email struct {
	Subject string   `json:"subject"`
	Body    Body     `json:"body"`
	From    Namepair `json:"from"`
	To      Namepair `json:"to"`
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

// Namepair An email address and a name
type Namepair struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}
