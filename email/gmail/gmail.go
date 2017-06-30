package gmail

import (
	"github.com/davidbanham/notify/config"
	"github.com/davidbanham/notify/types"
	"net/smtp"
)

// Send an email via gmail
func Send(e types.Email) error {
	pass := config.EmailSMTPPass

	msg := "Return-Path: " + e.From.Address + "\n" +
		"From: " + e.From.Address + "\n" +
		"To: " + e.To.Address + "\n" +
		"Subject: " + e.Subject + "\n\n" +
		e.Body.Text

	return smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", e.From.Address, pass, "smtp.gmail.com"),
		e.From.Address, []string{e.To.Address}, []byte(msg))
}
