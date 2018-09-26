package helper

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func GetAWSConfig(accessKeyID string, secretAccessKey string, bucket string, region string) (*aws.Config, error) {
	creds := credentials.NewStaticCredentials(accessKeyID, secretAccessKey, "")
	var err error
	if len(region) == 0 {
		sess := session.Must(session.NewSession(&aws.Config{
			Credentials: creds,
		}))
		region, err = s3manager.GetBucketRegion(context.Background(), sess, bucket, endpoints.UsEast1RegionID)
		fmt.Println("region:", region)
		if err != nil {
			return nil, err
		}
	}
	config := &aws.Config{
		Region:      aws.String(region),
		Credentials: creds,
	}
	return config, nil
}
