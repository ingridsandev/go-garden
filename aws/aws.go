package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Aws func
func Aws(region string) *session.Session {
	profile := "Isantos"
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String(region)},
		Profile: profile,
	}))

	return sess
}
