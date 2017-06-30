package mandrill

import (
	"github.com/davidbanham/notify/config"
	"github.com/davidbanham/notify/types"
	m "github.com/keighl/mandrill"
)

var client *m.Client

func init() {
	client = m.ClientWithKey(config.EmailMandrillKey)
}

// Send an email via Mandrill
func Send(e types.Email) error {
	message := &m.Message{}
	message.AddRecipient(e.To, e.ToName, "to")
	message.FromEmail = config.EmailFrom
	message.FromName = config.EmailFromName
	message.Subject = e.Subject
	if e.Body.HTML != "" {
		message.HTML = e.Body.HTML
	} else {
		message.HTML = e.Body.Text
	}
	message.Text = e.Body.Text

	_, err := client.MessagesSend(message)
	return err
}
