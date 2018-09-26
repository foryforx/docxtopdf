package usecase_test

import (
	"errors"
	"testing"

	"github.com/karuppaiah/docxtopdf/external/awsconn/mocks"
	"github.com/karuppaiah/docxtopdf/external/awsconn/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPullFileFromS3Success(t *testing.T) {
	mockRepo := new(mocks.IRepository)

	mockRepo.On("PullFileFromS3", mock.AnythingOfType("string"), mock.Anything, mock.Anything).Return(nil)

	u := usecase.NewEUseCase(mockRepo)

	err := u.PullFileFromS3("bucket", "destfile", "sourcefile")

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)

}

func TestPullFileFromS3Failure(t *testing.T) {
	mockRepo := new(mocks.IRepository)

	mockRepo.On("PullFileFromS3", mock.AnythingOfType("string"), mock.Anything, mock.Anything).Return(errors.New("Unexpexted Error"))

	u := usecase.NewEUseCase(mockRepo)

	err := u.PullFileFromS3("bucket", "destfile", "sourcefile")

	assert.Error(t, err)

	mockRepo.AssertExpectations(t)

}

func TestPushFileToS3Success(t *testing.T) {
	mockRepo := new(mocks.IRepository)

	mockRepo.On("PutFileToS3", mock.AnythingOfType("string"), mock.Anything, mock.Anything).Return(nil)

	u := usecase.NewEUseCase(mockRepo)

	err := u.PutFileToS3("bucket", "destfile", "sourcefile")

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)

}

func TestPushFileToS3Failure(t *testing.T) {
	mockRepo := new(mocks.IRepository)

	mockRepo.On("PutFileToS3", mock.AnythingOfType("string"), mock.Anything, mock.Anything).Return(errors.New("Unexpexted Error"))

	u := usecase.NewEUseCase(mockRepo)

	err := u.PutFileToS3("bucket", "destfile", "sourcefile")

	assert.Error(t, err)

	mockRepo.AssertExpectations(t)

}

func TestDelFileInS3Success(t *testing.T) {
	mockRepo := new(mocks.IRepository)

	mockRepo.On("DeleteFileFromS3", mock.AnythingOfType("string"), mock.Anything).Return(nil)

	u := usecase.NewEUseCase(mockRepo)

	err := u.DeleteFileFromS3("bucket", "destfile")

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)

}

func TestDelFileInS3Failure(t *testing.T) {
	mockRepo := new(mocks.IRepository)

	mockRepo.On("DeleteFileFromS3", mock.AnythingOfType("string"), mock.Anything).Return(errors.New("Unexpexted Error"))

	u := usecase.NewEUseCase(mockRepo)

	err := u.DeleteFileFromS3("bucket", "destfile")

	assert.Error(t, err)

	mockRepo.AssertExpectations(t)

}
