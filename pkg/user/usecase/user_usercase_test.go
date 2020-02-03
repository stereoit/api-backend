package usecase

import (
	"errors"
	"testing"

	mocksRepo "github.com/stereoit/eventival/mocks/user/domain/repository"
	mocksService "github.com/stereoit/eventival/mocks/user/domain/service"

	"github.com/stereoit/eventival/pkg/user/domain/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_NewUserUsecase(t *testing.T) {
	assert := assert.New(t)

	mockRepo := &mocksRepo.UserRepository{}
	mockService := &mocksService.UserService{}
	usecase := NewUserUsecase(mockRepo, mockService)
	assert.NotNil(usecase, "NewUserUsecase should return an instance")
}

func Test_ListAllUsers(t *testing.T) {
	assert := assert.New(t)

	// handle repository error
	mockRepo := &mocksRepo.UserRepository{}
	mockRepo.On("FindAll").Return(nil, errors.New("repo error")).Once()
	mockService := &mocksService.UserService{}
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
	mockRepo := &mocksRepo.UserRepository{}
	mockService := &mocksService.UserService{}
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

func Test_RegisterUser_Duplicated(t *testing.T) {
	assert := assert.New(t)
	email := "user@example.com"
	mockRepo := &mocksRepo.UserRepository{}
	mockService := &mocksService.UserService{}
	usecase := NewUserUsecase(mockRepo, mockService)

	// handle error of service
	mockService.On("Duplicated", email).Return(errors.New("service error")).Once()
	_, err := usecase.RegisterUser(email)
	assert.NotNil(err, "service Duplicated might throw error")
	mockService.AssertExpectations(t)

}

func Test_RegisterUser_RepoError(t *testing.T) {
	assert := assert.New(t)
	email := "user@example.com"
	mockRepo := &mocksRepo.UserRepository{}
	mockService := &mocksService.UserService{}
	usecase := NewUserUsecase(mockRepo, mockService)

	// handle error of repo
	mockService.On("Duplicated", email).Return(nil).Once()
	mockRepo.On("Save", mock.Anything).Return(errors.New("repo error"))
	_, err := usecase.RegisterUser(email)
	assert.NotNil(err, "repo Save might throw error")
	mockService.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func Test_RegisterUser(t *testing.T) {
	assert := assert.New(t)
	email := "user@example.com"
	mockRepo := &mocksRepo.UserRepository{}
	mockService := &mocksService.UserService{}
	usecase := NewUserUsecase(mockRepo, mockService)

	// handle success of registration
	mockService.On("Duplicated", email).Return(nil).Once()
	mockRepo.On("Save", mock.Anything).Return(nil).Once()
	id, _ := usecase.RegisterUser(email)
	assert.NotNil(id, "Registeruser should return ID of new user")
	mockService.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func Test_UpdateUser(t *testing.T) {
	assert := assert.New(t)
	mockRepo := &mocksRepo.UserRepository{}
	mockService := &mocksService.UserService{}
	usecase := NewUserUsecase(mockRepo, mockService)
	user := &User{
		ID: "8e492e66-9af8-48ab-a22d-61cbaaf333fa",
		// Email:     "user@example.com",
		FirstName: "Václav",
		LastName:  "Havel",
	}

	// test non valid user
	err := usecase.UpdateUser(user)
	assert.NotNil(err, "non valid user should have error")

	// test repository error
	user.Email = "user@example.com"
	mockRepo.On("Save", mock.Anything).Return(errors.New("repo error")).Once()
	err = usecase.UpdateUser(user)
	assert.NotNil(err, "handling repository error")
	mockRepo.AssertExpectations(t)

	// test valid user
	mockRepo.On("Save", mock.Anything).Return(nil).Once()
	err = usecase.UpdateUser(user)
	assert.Nil(err, "updating user should be OK")
	mockRepo.AssertExpectations(t)
}
