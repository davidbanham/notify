package sms

import (
	"errors"
	"github.com/davidbanham/notify/config"
	"github.com/davidbanham/notify/sms/amazon"
	"github.com/davidbanham/notify/types"
	"log"
)

var sender func(types.SMS) error

func init() {
	provider := config.EmailProvider

	providers := map[string]bool{
		"amazon": true,
		"none":   true,
		"test":   true,
	}

	if !providers[provider] {
		log.Fatal("Invalid sms provider specified ", provider, " valid providers are ", providers)
	}

	switch provider {
	case "amazon":
		sender = amazon.Send
		return
	default:
		sender = invalid
		return
	}
}

func invalid(e types.SMS) error {
	return errors.New("No valid sms provider configured")
}

func test(e types.SMS) error {
	return nil
}

// Send an SMS via the configured provider
func Send(e types.SMS) error {
	if e.From == "" {
		e.From = config.SmsFrom
	}
	return sender(e)
}
