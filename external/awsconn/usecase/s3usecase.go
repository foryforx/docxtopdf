package usecase

import (
	"fmt"

	"github.com/karuppaiah/docxtopdf/external/awsconn"
)

//eUseCase with Connection to AWS
type eUsecase struct {
	eRepos awsconn.IRepository
	// contextTimeout time.Duration
}

// To create new UseCase
func NewEUseCase(a awsconn.IRepository) awsconn.IUsecase {
	return &eUsecase{a}
}

func (a *eUsecase) PullFileFromS3(bucket string, destFilename string, sourceFilename string) error {
	fmt.Println("Usecase|in PullFileFromS3")
	err := a.eRepos.PullFileFromS3(bucket, destFilename, sourceFilename)
	if err != nil {
		return err
	}

	return nil

}

func (a *eUsecase) PutFileToS3(bucket string, destFilename string, sourceFilename string) error {
	fmt.Println("Usecase|in PutFileToS3")

	err := a.eRepos.PutFileToS3(bucket, destFilename, sourceFilename)
	if err != nil {
		return err
	}

	return nil

}

func (a *eUsecase) DeleteFileFromS3(bucket string, destFilename string) error {
	fmt.Println("Usecase|in DelFileInS3")

	err := a.eRepos.DeleteFileFromS3(bucket, destFilename)
	if err != nil {
		return err
	}

	return nil

}
