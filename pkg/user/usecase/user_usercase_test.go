package usecase

import (
	"errors"
	"github.com/stereoit/eventival/pkg/mocks"
	"github.com/stereoit/eventival/pkg/user/domain/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_NewUserUsecase(t *testing.T) {
	assert := assert.New(t)

	mockRepo := &mocks.UserRepository{}
	mockService := &mocks.UserService{}
	usecase := NewUserUsecase(mockRepo, mockService)
	assert.NotNil(usecase, "NewUserUsecase should return an instance")
}

func Test_ListAllUsers(t *testing.T) {
	assert := assert.New(t)

	// handle repository error
	mockRepo := &mocks.UserRepository{}
	mockRepo.On("FindAll").Return(nil, errors.New("repo error")).Once()
	mockService := &mocks.UserService{}
	usecase := NewUserUsecase(mockRepo, mockService)
	_, err := usecase.ListAllUsers()
	assert.NotNil(err, "ListAllUsers repo might throw error")
	mockRepo.AssertExpectations(t)

	usersResult := []*model.User{&model.User{}, &model.User{}}
	mockRepo.On("FindAll").Return(usersResult, nil).Once()
	users, _ := usecase.ListAllUsers()
	assert.NotNil(users, "ListAllUsers should return some users")
	mockRepo.AssertExpectations(t)
}

func Test_FindByID(t *testing.T) {
	assert := assert.New(t)
	id := "1"
	mockRepo := &mocks.UserRepository{}
	mockService := &mocks.UserService{}
	usecase := NewUserUsecase(mockRepo, mockService)

	// handle repository error
	mockRepo.On("FindByID", id).Return(nil, errors.New("repo error")).Once()
	_, err := usecase.FindByID(id)
	assert.NotNil(err, "FindByID repo might throw error")
	mockRepo.AssertExpectations(t)

	// handle found user
	mockRepo.On("FindByID", id).Return(&model.User{}, nil).Once()
	user, _ := usecase.FindByID(id)
	assert.NotNil(user, "FindByID should return user by ID")
	mockRepo.AssertExpectations(t)
}

func Test_RegisterUser(t *testing.T) {
	assert := assert.New(t)
	email := "user@example.com"
	mockRepo := &mocks.UserRepository{}
	mockService := &mocks.UserService{}
	usecase := NewUserUsecase(mockRepo, mockService)

	// handle error of service
	mockService.On("Duplicated", email).Return(errors.New("service error")).Once()
	_, err := usecase.RegisterUser(email)
	assert.NotNil(err, "service Duplicated might throw error")
	mockService.AssertExpectations(t)

	// handle error of repo
	mockService.On("Duplicated", email).Return(nil).Once()
	mockRepo.On("Save", mock.Anything).Return(errors.New("repo error"))
	_, err = usecase.RegisterUser(email)
	assert.NotNil(err, "repo Save might throw error")
	mockService.AssertExpectations(t)
	mockRepo.AssertExpectations(t)

	// handle success of registration
	mockService.On("Duplicated", email).Return(nil).Once()
	mockRepo.On("Save", mock.Anything).Return(nil)
	id, _ := usecase.RegisterUser(email)
	assert.NotNil(id, "Registeruser should return ID of new user")
	mockService.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}
