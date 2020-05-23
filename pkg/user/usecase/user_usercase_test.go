package usecase

import (
	"errors"
	"testing"

	mocks "github.com/stereoit/eventival/mocks/user/domain/repository"
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
	page := 0
	limit := 1

	// handle repository error
	mockRepo := &mocksRepo.UserRepository{}
	mockRepo.On("FindAll", page, limit).Return(nil, errors.New("repo error")).Once()
	mockService := &mocksService.UserService{}
	usecase := NewUserUsecase(mockRepo, mockService)
	_, err := usecase.ListAllUsers(page, limit)
	assert.NotNil(err, "ListAllUsers repo might throw error")
	mockRepo.AssertExpectations(t)

	usersResult := []*model.User{&model.User{}, &model.User{}}
	mockRepo.On("FindAll", page, limit).Return(usersResult, nil).Once()
	users, _ := usecase.ListAllUsers(page, limit)
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

func Test_RegisterUser(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(*mocksRepo.UserRepository, *mocksService.UserService)
		wantErr bool
		email   string
	}{
		{
			name: "OK",
			mock: func(mockRepo *mocks.UserRepository, mockService *mocksService.UserService) {
				mockService.On("Duplicated", mock.Anything).Return(nil).Once()
				mockRepo.On("Save", mock.Anything).Return(nil)
			},
			wantErr: false,
			email:   "user@example.com",
		},
		{
			name: "Duplicated",
			mock: func(mockRepo *mocks.UserRepository, mockService *mocksService.UserService) {
				mockService.On("Duplicated", mock.Anything).Return(errors.New("Duplicated email")).Once()
			},
			wantErr: true,
			email:   "user@example.com",
		},
		{
			name: "Bad email",
			mock: func(mockRepo *mocks.UserRepository, mockService *mocksService.UserService) {
				mockService.On("Duplicated", mock.Anything).Return(nil).Once()
			},
			wantErr: true,
			email:   "invalid-email",
		},
		{
			name: "Repo error",
			mock: func(mockRepo *mocks.UserRepository, mockService *mocksService.UserService) {
				mockService.On("Duplicated", mock.Anything).Return(nil).Once()
				mockRepo.On("Save", mock.Anything).Return(errors.New("Repo error"))
			},
			wantErr: true,
			email:   "user@example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocksRepo.UserRepository{}
			mockService := &mocksService.UserService{}
			usecase := NewUserUsecase(mockRepo, mockService)
			tt.mock(mockRepo, mockService)
			_, err := usecase.RegisterUser(tt.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.Delete() error = %v , wantErr %v", err, tt.wantErr)
				return
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func Test_UpdateUser(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(*mocks.UserRepository)
		wantErr bool
		user    *User
	}{
		{
			name: "OK",
			mock: func(mockRepo *mocks.UserRepository) {
				mockRepo.On("Update", mock.Anything).Return(nil)
			},
			wantErr: false,
			user: &User{
				ID:        "8e492e66-9af8-48ab-a22d-61cbaaf333fa",
				Email:     "user@example.com",
				FirstName: "Václav",
				LastName:  "Havel",
			},
		},
		{
			name:    "Non valid user",
			mock:    func(*mocks.UserRepository) {},
			wantErr: true,
			user:    &User{},
		},
		{
			name: "Repo error",
			mock: func(mockRepo *mocks.UserRepository) {
				mockRepo.On("Update", mock.Anything).Return(errors.New("Repo error"))
			},
			wantErr: true,
			user: &User{
				ID:    "8e492e66-9af8-48ab-a22d-61cbaaf333fa",
				Email: "user@example.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocksRepo.UserRepository{}
			mockService := &mocksService.UserService{}
			usecase := NewUserUsecase(mockRepo, mockService)
			tt.mock(mockRepo)
			err := usecase.UpdateUser(tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.Delete() error = %v , wantErr %v", err, tt.wantErr)
				return
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	userID := "8e492e66-9af8-48ab-a22d-61cbaaf333fa"
	tests := []struct {
		name    string
		mock    func(*mocks.UserRepository)
		wantErr bool
	}{
		{
			name:    "Delete User",
			wantErr: false,
			mock: func(mockRepo *mocks.UserRepository) {
				mockRepo.On("Delete", mock.Anything).Return(nil)
			},
		},
		{
			name: "Repository error",
			mock: func(mockRepo *mocks.UserRepository) {
				mockRepo.On("Delete", mock.Anything).Return(errors.New("Repo error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocksRepo.UserRepository{}
			mockService := &mocksService.UserService{}
			usecase := NewUserUsecase(mockRepo, mockService)
			tt.mock(mockRepo)
			err := usecase.DeleteUser(userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.Delete() error = %v , wantErr %v", err, tt.wantErr)
				return
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
