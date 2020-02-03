package service

import (
	"errors"
	"testing"

	mocks "github.com/stereoit/eventival/mocks/user/domain/repository"

	"github.com/stereoit/eventival/pkg/user/domain/model"
	"github.com/stretchr/testify/assert"
)

func Test_NewUserService(t *testing.T) {
	assert := assert.New(t)
	mockRepo := &mocks.UserRepository{}

	service := NewUserService(mockRepo)
	assert.NotNil(service, "NewUserService should return an intance")
}

func Test_Duplicated(t *testing.T) {
	assert := assert.New(t)
	mockRepo := &mocks.UserRepository{}
	testEmail := "test-email@example.com"

	// no error for unique email
	service := NewUserService(mockRepo)
	mockRepo.On("FindByEmail", testEmail).Return(nil, nil).Once()
	err := service.Duplicated(testEmail)
	assert.Nil(err, "unique email should return no error")
	mockRepo.AssertExpectations(t)

	// error for duplicated user
	mockRepo.On("FindByEmail", testEmail).Return(&model.User{}, nil).Once()
	err = service.Duplicated(testEmail)
	assert.NotNil(err, "duplicated email returns error")
	mockRepo.AssertExpectations(t)

	// error for repository error
	mockRepo.On("FindByEmail", testEmail).Return(nil, errors.New("error")).Once()
	err = service.Duplicated(testEmail)
	assert.NotNil(err, "repository might throw error as well")
	mockRepo.AssertExpectations(t)
}
