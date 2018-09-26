package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	awshelper "github.com/karuppaiah/docxtopdf/external/awsconn/helper"
	awsRepo "github.com/karuppaiah/docxtopdf/external/awsconn/repository"
	awsusecase "github.com/karuppaiah/docxtopdf/external/awsconn/usecase"
	pdfapi "github.com/karuppaiah/docxtopdf/pdfgenerate/api"
	pdfUsecase "github.com/karuppaiah/docxtopdf/pdfgenerate/usecase"
)

var (
	NWorkers = flag.Int("n", 1, "The number of workers to start")
	HTTPAddr = flag.String("http", "127.0.0.1:8000", "Address to listen for HTTP requests on")
)

func main1() {
	// Parse the command-line flags.
	flag.Parse()

	// Start the dispatcher.
	fmt.Println("Starting the dispatcher")
	pdfUsecase.StartDispatcher(*NWorkers)

	// Register our collector as an HTTP handler function.
	fmt.Println("Registering the collector")
	http.HandleFunc("/work", pdfapi.Collector)

	// Start the HTTP server!
	fmt.Println("HTTP server listening on", *HTTPAddr)
	if err := http.ListenAndServe(*HTTPAddr, nil); err != nil {
		fmt.Println(err.Error())
	}
}
func main() {

	bucket := os.Getenv("AWS_REPORTS_BUCKET")
	keyVal := os.Getenv("AWS_ACCESS_KEY_ID")
	secretVal := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")

	config, err := awshelper.GetAWSConfig(keyVal, secretVal, bucket, region)
	if err != nil {
		fmt.Println("Unable to create config for s3conn: ", err)

	}
	sess := session.New(config)
	svc := s3.New(sess)
	s3Repository := awsRepo.NewERepository(config, svc)
	s3Usecase := awsusecase.NewEUseCase(s3Repository)
	// Push file to s3
	destFilename := "sample0.docx" + "_" + "kal"
	sourceLocalFilename := "sample0.docx"
	err = s3Usecase.PutFileToS3(bucket, destFilename, sourceLocalFilename)
	if err != nil {
		fmt.Println("Unable to Put file in s3conn: ", err)
	}
	// Pull file from S3
	destFilename = "pull_sample0.docx" + "_" + "kal"
	sourceLocalFilename = "sample0.docx" + "_" + "kal"
	err = s3Usecase.PullFileFromS3(bucket, destFilename, sourceLocalFilename)
	if err != nil {
		fmt.Println("Unable to pull file in s3conn: ", err)
	}

	// Delete file from S3
	destFilename = "sample0.docx" + "_" + "kal"

	err = s3Usecase.DeleteFileFromS3(bucket, destFilename)
	if err != nil {
		fmt.Println("Unable to Delete file in s3conn: ", err)
	}

}
