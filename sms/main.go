package sms

import (
	"errors"
	"github.com/davidbanham/notify/sms/amazon"
	"github.com/davidbanham/notify/types"
	"os"
)

func init() {
}

func Send(m types.SMS) error {
	switch os.Getenv("NOTIFY_SMS_PROVIDER") {
	case "amazon":
		return amazonSms.Send(m)
	default:
		return errors.New("No valid sms provider configured")
	}
}
