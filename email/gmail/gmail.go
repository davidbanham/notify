package gmail

import (
	"github.com/davidbanham/notify/config"
	"github.com/davidbanham/notify/types"
	"net/smtp"
)

// Send an email via gmail
func Send(e types.Email) error {
	from := config.EmailFrom
	pass := config.EmailSMTPPass

	msg := "Return-Path: " + from + "\n" +
		"From: " + e.From + "\n" +
		"To: " + e.To + "\n" +
		"Subject: " + e.Subject + "\n\n" +
		e.Body.Text

	return smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{e.To}, []byte(msg))
}
