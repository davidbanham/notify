package config

import (
	"github.com/davidbanham/required_env"
	"os"
)

// Port The port to listen on
var Port string

// EmailProvider The provider to use for email notifications
var EmailProvider string

// SmsProvider The provider to use for sms notifications
var SmsProvider string

// SmsFrom The default From for an SMS
var SmsFrom string

// EmailSMTPPass The password to use if sending email via SMTP/Gmail
var EmailSMTPPass string

// EmailFrom The default from: address for emails sent
var EmailFrom string

// EmailFromName The default name that email should appear to come from
var EmailFromName string

// EmailMandrillKey The key to use if sending email via Mandrill
var EmailMandrillKey string

// AwsAccessKeyID your AWS access key ID
var AwsAccessKeyID string

// AwsSecretAccessKey your AWS access key
var AwsSecretAccessKey string

// AwsRegion your AWS access region
var AwsRegion string

// Testing Whether we are in test mode
var Testing string

func init() {
	required_env.Ensure(map[string]string{
		"PORT":                  "",
		"NOTIFY_EMAIL_PROVIDER": "",
		"NOTIFY_SMS_PROVIDER":   "",
		"NOTIFY_SMS_FROM":       "",
	})

	switch os.Getenv("NOTIFY_EMAIL_PROVIDER") {
	case "gmail":
		required_env.Ensure(map[string]string{
			"NOTIFY_EMAIL_SMTP_PASS": "",
			"NOTIFY_EMAIL_FROM":      "",
		})
	case "mandrill":
		required_env.Ensure(map[string]string{
			"NOTIFY_EMAIL_MANDRILL_KEY": "",
			"NOTIFY_EMAIL_FROM":         "",
		})
	}
	switch os.Getenv("NOTIFY_SMS_PROVIDER") {
	case "amazon":
		required_env.Ensure(map[string]string{
			"AWS_ACCESS_KEY_ID":     "",
			"AWS_SECRET_ACCESS_KEY": "",
			"AWS_REGION":            "",
		})
	}

	Port = os.Getenv("PORT")
	EmailProvider = os.Getenv("NOTIFY_EMAIL_PROVIDER")
	SmsProvider = os.Getenv("NOTIFY_SMS_PROVIDER")
	SmsFrom = os.Getenv("NOTIFY_SMS_FROM")
	EmailSMTPPass = os.Getenv("NOTIFY_EMAIL_SMTP_PASS")
	EmailFrom = os.Getenv("NOTIFY_EMAIL_FROM")
	EmailFromName = os.Getenv("NOTIFY_EMAIL_FROM_NAME")
	AwsAccessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
	AwsSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	AwsRegion = os.Getenv("AWS_REGION")
	Testing = os.Getenv("TESTING")

	if EmailFromName == "" {
		EmailFromName = EmailFrom
	}
}
