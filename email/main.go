package email

import (
	"errors"
	"fmt"
	"github.com/davidbanham/notify/config"
	"github.com/davidbanham/notify/email/amazon"
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
		"amazon":   true,
		"none":     true,
		"test":     true,
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
	case "amazon":
		sender = amazon.SendFactory(amazon.Init())
		return
	case "test":
		sender = test
		return
	default:
		sender = invalid
		return
	}
}

func invalid(e types.Email) error {
	return errors.New("No valid email provider configured")
}

func test(e types.Email) error {
	fmt.Println("Subject", e.Subject)
	fmt.Println("From", e.From)
	fmt.Println("To", e.To)
	fmt.Println("Body", e.Body)
	fmt.Println("Recieved email. Dropping it due to test route being configured.")
	if e.To.Address == "" {
		return errors.New("No address given to send email to")
	}
	return nil
}

// Send an email via the configured provider
func Send(e types.Email) (types.Email, error) {
	if e.From.Name == "" {
		e.From.Name = config.EmailFromName
	}

	if e.From.Address == "" {
		e.From.Address = config.EmailFrom
	}

	return e, sender(e)
}
