package amazon

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/davidbanham/notify/config"
	"github.com/davidbanham/notify/types"
)

func Send(m types.SMS) error {
	svc := sns.New(
		session.New(),
		&aws.Config{
			Region: aws.String(config.AWS_REGION),
		},
	)

	params := &sns.PublishInput{
		Message: aws.String(m.Body),
		MessageAttributes: map[string]*sns.MessageAttributeValue{
			"AWS.SNS.SMS.SenderID": {
				DataType:    aws.String("String"),
				StringValue: aws.String(m.From),
			},
		},
		PhoneNumber: aws.String(m.To),
	}

	_, err := svc.Publish(params)
	return err
}
