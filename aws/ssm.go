package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func PutParameter(region string, name string, value string) string {
	ssmService := ssm.New(Aws(region))

	_, err := ssmService.PutParameter(&ssm.PutParameterInput{
		Name:           aws.String(name),
		Type:           aws.String("String"),
		Value:          aws.String(value),
	})

	if err != nil {
		return "Something went wrong while putting parameter into ssm"
	}

	return "OK"
}

func GetParameter(region string, name string) string {
	ssmService := ssm.New(Aws(region))

	response, err := ssmService.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(name),
	})

	if err != nil {
		return "Something went wrong while getting parameter from ssm"
	}

	return *response.Parameter.Value
}

func DeleteParameter(region string, name string) string {
	ssmService := ssm.New(Aws(region))

	_, err := ssmService.DeleteParameter(&ssm.DeleteParameterInput{
		Name:           aws.String(name),
	})

	if err != nil {
		return "Something went wrong while deleting parameter from ssm"
	}

	return "OK"
}
