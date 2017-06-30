package config

import (
	"github.com/davidbanham/required_env"
	"os"
)

var PORT string
var NOTIFY_EMAIL_PROVIDER string
var NOTIFY_SMS_PROVIDER string
var NOTIFY_EMAIL_SMTP_PASS string
var NOTIFY_EMAIL_FROM string
var AWS_ACCESS_KEY_ID string
var AWS_SECRET_ACCESS_KEY string
var AWS_REGION string
var TESTING string

func init() {
	required_env.Ensure(map[string]string{
		"PORT":                  "",
		"NOTIFY_EMAIL_PROVIDER": "",
		"NOTIFY_SMS_PROVIDER":   "",
	})

	switch os.Getenv("NOTIFY_EMAIL_PROVIDER") {
	case "gmail":
		required_env.Ensure(map[string]string{
			"NOTIFY_EMAIL_SMTP_PASS": "",
			"NOTIFY_EMAIL_FROM":      "",
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

	PORT = os.Getenv("PORT")
	NOTIFY_EMAIL_PROVIDER = os.Getenv("NOTIFY_EMAIL_PROVIDER")
	NOTIFY_SMS_PROVIDER = os.Getenv("NOTIFY_SMS_PROVIDER")
	NOTIFY_EMAIL_SMTP_PASS = os.Getenv("NOTIFY_EMAIL_SMTP_PASS")
	NOTIFY_EMAIL_FROM = os.Getenv("NOTIFY_EMAIL_FROM")
	AWS_ACCESS_KEY_ID = os.Getenv("AWS_ACCESS_KEY_ID")
	AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")
	AWS_REGION = os.Getenv("AWS_REGION")
	TESTING = os.Getenv("TESTING")
}
