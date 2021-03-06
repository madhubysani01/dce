package rolemanager

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/iam"
)

func isAWSAlreadyExistsError(err error) bool {
	aerr, ok := err.(awserr.Error)
	if ok {
		switch aerr.Code() {
		case iam.ErrCodeEntityAlreadyExistsException:
			return true
		}
	}

	return false
}
