package repository

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/karuppaiah/docxtopdf/external/awsconn"
)

//eUseCase with Connection to AWS
type eRepository struct {
	Conn *aws.Config
	Svc  *s3.S3 //*s3iface.S3API
}

// To create new UseCase
func NewERepository(conn *aws.Config, svc *s3.S3) awsconn.IRepository {
	// sess := session.New(conn)
	// svc := s3.New(sess)
	return &eRepository{Conn: conn, Svc: svc}
}

func (a *eRepository) PullFileFromS3(bucket string, destFilename string, sourceFilename string) error {
	fmt.Println("Repo: in PullFileToS3")

	sess := session.New(a.Conn)

	downloader := s3manager.NewDownloader(sess)
	f, err := os.Create(destFilename)
	fmt.Println("Creating file", destFilename)
	if err != nil {
		return err
	}
	defer f.Close()

	n, pullErr := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(sourceFilename),
	})
	fmt.Printf("file downloaded, %d bytes\n", n)
	return pullErr

	// sess := session.New(config)
	// svc := s3.New(sess)
	// input := &s3.GetObjectInput{
	// 	Bucket: aws.String(bucket),
	// 	Key:    aws.String(sourceFilename),
	// }
	// fmt.Println(sourceFilename)
	// result, err := svc.GetObject(input)
	// if err != nil {
	// 	if aerr, ok := err.(awserr.Error); ok {
	// 		switch aerr.Code() {
	// 		case s3.ErrCodeNoSuchKey:
	// 			fmt.Println(s3.ErrCodeNoSuchKey, aerr.Error())
	// 		default:
	// 			fmt.Println(aerr.Error())
	// 		}
	// 	} else {
	// 		// Print the error, cast err to awserr.Error to get the Code and
	// 		// Message from an error.
	// 		fmt.Println(err.Error())
	// 	}
	// 	return nil
	// }

	// fmt.Println(result)
	// return nil
}

func (a *eRepository) PutFileToS3(bucket string, destFilename string, sourceFilename string) error {
	fmt.Println("in PutFileToS3")

	var timeout = 100 * time.Second

	ctx, cancelFn := context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	f, err := os.Open(sourceFilename)
	fmt.Println("Opening file", sourceFilename)
	if err != nil {
		return nil
	}
	defer f.Close()
	_, putErr := a.Svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(destFilename),
		Body:   f,
	})
	fmt.Println("done with put s3")
	return putErr
}

func (a *eRepository) DeleteFileFromS3(bucket string, destFilename string) error {
	fmt.Println("in DeleteFileFromS3")

	input := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(destFilename),
	}

	result, err := a.Svc.DeleteObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
				return aerr
			}
		}
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())

		return err
	}

	fmt.Println(result)

	return nil
}
