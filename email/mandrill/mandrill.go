package mandrill

import (
	"github.com/davidbanham/notify/config"
	"github.com/davidbanham/notify/types"
	m "github.com/keighl/mandrill"
)

var client *m.Client

func init() {
	client = m.ClientWithKey(config.NOTIFY_EMAIL_MANDRILL_KEY)
}

func Send(e types.Email) error {
	message := &m.Message{}
	message.AddRecipient(e.To, e.To, "to")
	message.FromEmail = config.NOTIFY_EMAIL_FROM
	message.FromName = config.NOTIFY_EMAIL_FROM
	message.Subject = e.Subject
	message.HTML = e.Body.Html
	message.Text = e.Body.Text

	_, err := client.MessagesSend(message)
	return err
}
