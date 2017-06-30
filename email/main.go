package email

import (
	"errors"
	"github.com/davidbanham/notify/config"
	"github.com/davidbanham/notify/email/gmail"
	"github.com/davidbanham/notify/email/mandrill"
	"github.com/davidbanham/notify/types"
	"log"
)

var sender func(types.Email) error

func init() {
	provider := config.EmailProvider

	providers := map[string]bool{
		"gmail":    true,
		"mandrill": true,
		"none":     true,
	}

	if !providers[provider] {
		log.Fatal("Invalid email provider specified ", provider, " valid providers are ", providers)
	}

	switch provider {
	case "gmail":
		sender = gmail.Send
		return
	case "mandrill":
		sender = mandrill.Send
		return
	default:
		sender = invalid
		return
	}
}

func invalid(e types.Email) error {
	return errors.New("No valid email provider configured")
}

// Send an email via the configured provider
func Send(e types.Email) error {
	if e.From.Name == "" {
		e.From.Name = config.EmailFromName
	}

	if e.From.Address == "" {
		e.From.Address = config.EmailFrom
	}

	return sender(e)
}
