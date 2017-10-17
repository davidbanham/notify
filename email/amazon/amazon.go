package amazon

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/davidbanham/notify/config"
	"github.com/davidbanham/notify/types"
)

func Init() *ses.SES {
	svc := ses.New(
		session.New(),
		&aws.Config{
			Region: aws.String(config.AwsRegion),
		},
	)
	return svc
}

func SendFactory(svc *ses.SES) types.EmailSender {
	return func(email types.Email) (err error) {
		body := &ses.Body{
			Text: &ses.Content{
				Data:    aws.String(email.Body.Text),
				Charset: aws.String("UTF8"),
			},
		}

		if email.Body.HTML != "" {
			body.Html = &ses.Content{
				Data:    aws.String(email.Body.HTML),
				Charset: aws.String("UTF8"),
			}
		}

		params := &ses.SendEmailInput{
			Destination: &ses.Destination{
				ToAddresses: []*string{
					aws.String(email.To.Address),
				},
			},
			Message: &ses.Message{
				Body: body,
				Subject: &ses.Content{
					Data:    aws.String(email.Subject),
					Charset: aws.String("UTF8"),
				},
			},
			Source: aws.String(email.From.Address),
		}
		_, err = svc.SendEmail(params)

		return
	}
}
