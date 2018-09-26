package awsconn

type IRepository interface {
	PullFileFromS3(bucket string, destFilename string, sourceFilename string) error
	PutFileToS3(bucket string, destFilename string, sourceFilename string) error
	DeleteFileFromS3(bucket string, destFilename string) error
}
