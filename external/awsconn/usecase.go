package awsconn

// Interface for Usecase. Can be used for testing if called from API layer
type IUsecase interface {
	PullFileFromS3(bucket string, destFilename string, sourceFilename string) error
	PutFileToS3(bucket string, destFilename string, sourceFilename string) error
	DeleteFileFromS3(bucket string, destFilename string) error
}
