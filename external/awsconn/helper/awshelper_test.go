package helper_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"

	awshelper "github.com/karuppaiah/docxtopdf/external/awsconn/helper"
	"github.com/stretchr/testify/assert"
)

func TestGetAWSConfigSuccess(t *testing.T) {
	creds := credentials.NewStaticCredentials("key", "secret", "")
	config, err := awshelper.GetAWSConfig("key", "secret", "bucket", "region")
	assert.Equal(t, aws.String("region"), config.Region)
	assert.Equal(t, creds, config.Credentials)
	assert.NoError(t, err)

}

func TestGetAWSConfigFailure(t *testing.T) {
	// creds := credentials.NewStaticCredentials("key", "secret", "")
	config, err := awshelper.GetAWSConfig("key", "secret", "bucket", "")
	fmt.Println(err)
	assert.Equal(t, aws.String("us-east-1"), config.Region)

	assert.NoError(t, err)

}
