package model

type S3Config struct {
	AccessKeyID      string
	SecretKey        string
	S3ReportBucket   string
	S3ImageBucket    string
	S3UserFileBucket string
	Region           string
}

func NewS3Config(accessKeyID string, secretKey string, reportBucket string, imageBucket string, fileBucket string, region string) *S3Config {
	return &S3Config{
		AccessKeyID:      accessKeyID,
		SecretKey:        secretKey,
		S3ReportBucket:   reportBucket,
		S3ImageBucket:    imageBucket,
		S3UserFileBucket: fileBucket,
		Region:           region,
	}
}
