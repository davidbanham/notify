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
	provider := config.NOTIFY_EMAIL_PROVIDER

	providers := map[string]bool{
		"amazon": true,
		"none":   true,
	}

	if !providers[provider] {
		log.Fatal("Invalid sms provider specified", provider, "valid providers are", providers)
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

func Send(e types.SMS) error {
	return sender(e)
}
