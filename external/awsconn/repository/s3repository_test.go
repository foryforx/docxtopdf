package repository_test

import (
	"testing"
)

//TODO: Yet to find a mock for S3 and write unit testing
func TestPutFileToS3Success(t *testing.T) {
	// mockS3 := new(mocks.IS3External)
	// config, err := awshelper.GetAWSConfig("key", "secretVal", "bucket", "region")
	// if err != nil {
	// 	fmt.Println("Unable to create config for s3conn: ", err)

	// }
	// mockS3.On("PutObjectWithContext", mock.AnythingOfType("Context"), mock.AnythingOfType("Interface")).Return(nil, nil)

	// u := repository.NewERepository(config, mockS3)

	// err = u.PutFileToS3("bucket", "destfile", "sourcefile")

	// assert.NoError(t, err)

	// mockS3.AssertExpectations(t)
}
