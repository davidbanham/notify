package email

import (
	"errors"
	"github.com/davidbanham/notify/email/gmail"
	"github.com/davidbanham/notify/types"
	"log"
	"os"
)

var sender func(types.Email) error

func init() {
	provider := os.Getenv("NOTIFY_EMAIL_PROVIDER")

	providers := map[string]bool{
		"gmail": true,
		"none":  true,
	}

	if !providers[provider] {
		log.Fatal("Invalid email provider specified", provider, "valid providers are", providers)
	}

	switch provider {
	case "gmail":
		sender = gmail.Send
		return
	default:
		sender = invalid
		return
	}
}

func invalid(e types.Email) error {
	return errors.New("No valid email provider configured")
}

func Send(e types.Email) error {
	return sender(e)
}
