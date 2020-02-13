package rest

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	mocks "github.com/stereoit/eventival/mocks/user/usecase"
	"github.com/stereoit/eventival/pkg/user/domain"
	"github.com/stereoit/eventival/pkg/user/usecase"

	"github.com/stretchr/testify/mock"
)

func Test_userService_RegisterUser(t *testing.T) {
	tests := []struct {
		name           string
		requestPayload string
		responseStatus int
		mock           func(*mocks.UserUsecase)
	}{
		{
			name: "OK",
			requestPayload: `
				{
					"email": "test@example.com",
					"firstName": "Test",
					"lastName": "User"
				}
				`,
			responseStatus: http.StatusCreated,
			mock: func(mockUserUsecase *mocks.UserUsecase) {
				mockUserUsecase.On("RegisterUser", mock.Anything).Return("1", nil)
				mockUserUsecase.On("FindByID", mock.Anything).Return(&usecase.User{}, nil)
				mockUserUsecase.On("UpdateUser", mock.Anything).Return(nil)
			},
		},
		{
			name:           "Duplicated user",
			requestPayload: "{\"email\": \"test@example.com\"}",
			responseStatus: http.StatusBadRequest,
			mock: func(mockUserUsecase *mocks.UserUsecase) {
				mockUserUsecase.On("RegisterUser", mock.Anything).Return("", &domain.DuplicateError{})
			},
		},
		{
			name:           "Missing expected attributes",
			requestPayload: "",
			responseStatus: http.StatusBadRequest,
			mock:           func(*mocks.UserUsecase) {},
		},
		{
			name: "Internal server error",
			requestPayload: `
				{
					"email": "test@example.com",
					"firstName": "Test",
					"lastName": "User"
				}
			`,
			responseStatus: http.StatusInternalServerError,
			mock: func(mockUserUsecase *mocks.UserUsecase) {
				mockUserUsecase.On("RegisterUser", mock.Anything).Return("", errors.New("Register user error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserUsecase := &mocks.UserUsecase{}
			service := &userService{mockUserUsecase}
			tt.mock(mockUserUsecase)

			req, _ := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(tt.requestPayload)))
			req.Header.Add("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(service.registerUser)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.responseStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.responseStatus)
			}
			mockUserUsecase.AssertExpectations(t)
		})
	}
}

func Test_userService_UpdateUser(t *testing.T) {
	tests := []struct {
		name           string
		requestPayload string
		responseStatus int
		mock           func(*mocks.UserUsecase)
	}{
		{
			name: "OK",
			requestPayload: `
				{
					"email": "test@example.com",
					"firstName": "Test",
					"lastName": "User"
				}
				`,
			responseStatus: http.StatusNoContent,
			mock: func(mockUserUsecase *mocks.UserUsecase) {
				user := &usecase.User{
					Email:     "test@example.com",
					ID:        "1",
					FirstName: "Test",
					LastName:  "User",
				}
				mockUserUsecase.On("FindByID", "1").Return(user, nil)
				mockUserUsecase.On("UpdateUser", mock.Anything).Return(nil)
			},
		},
		{
			name:           "User Not Found",
			requestPayload: "",
			responseStatus: http.StatusNotFound,
			mock: func(mockUserUsecase *mocks.UserUsecase) {
				mockUserUsecase.On("FindByID", mock.Anything).Return(nil, errors.New("FindByID user error"))
			},
		},
		{
			name:           "Invalid payload",
			requestPayload: "invalid format",
			responseStatus: http.StatusBadRequest,
			mock: func(mockUserUsecase *mocks.UserUsecase) {
				mockUserUsecase.On("FindByID", mock.Anything).Return(&usecase.User{}, nil)
			},
		},
		{
			name: "Internal server error",
			requestPayload: `
				{
					"email": "test@example.com",
					"firstName": "Test",
					"lastName": "User"
				}
			`,
			responseStatus: http.StatusInternalServerError,
			mock: func(mockUserUsecase *mocks.UserUsecase) {
				mockUserUsecase.On("FindByID", mock.Anything).Return(&usecase.User{}, nil)
				mockUserUsecase.On("UpdateUser", mock.Anything).Return(errors.New("UpdatedUser error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserUsecase := &mocks.UserUsecase{}
			service := &userService{mockUserUsecase}
			router := service.Routes()
			tt.mock(mockUserUsecase)

			req, _ := http.NewRequest("PATCH", "/1", bytes.NewBuffer([]byte(tt.requestPayload)))
			req.Header.Add("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req.WithContext(context.TODO()))
			// handler := http.HandlerFunc(service.updateUser)
			// handler.ServeHTTP(rr, req.WithContext(context.TODO()))

			if status := rr.Code; status != tt.responseStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.responseStatus)
			}
			mockUserUsecase.AssertExpectations(t)
		})
	}
}

func Test_userService_DeleteUser(t *testing.T) {
	tests := []struct {
		name           string
		responseStatus int
		mock           func(*mocks.UserUsecase)
	}{
		{
			name:           "OK",
			responseStatus: http.StatusNoContent,
			mock: func(mockUserUsecase *mocks.UserUsecase) {
				mockUserUsecase.On("DeleteUser", mock.Anything).Return(nil)
			},
		},
		{
			name:           "Usecase Error",
			responseStatus: http.StatusInternalServerError,
			mock: func(mockUserUsecase *mocks.UserUsecase) {
				mockUserUsecase.On("DeleteUser", mock.Anything).Return(errors.New("Usecase error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserUsecase := &mocks.UserUsecase{}
			service := &userService{mockUserUsecase}
			router := service.Routes()
			tt.mock(mockUserUsecase)

			req, _ := http.NewRequest("DELETE", "/1", nil)
			req.Header.Add("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req.WithContext(context.TODO()))

			if status := rr.Code; status != tt.responseStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.responseStatus)
			}
			mockUserUsecase.AssertExpectations(t)
		})
	}
}

func Test_userService_ListUser(t *testing.T) {
	tests := []struct {
		name           string
		responseStatus int
		mock           func(*mocks.UserUsecase)
	}{
		{
			name:           "OK",
			responseStatus: http.StatusOK,
			mock: func(mockUserUsecase *mocks.UserUsecase) {
				mockUserUsecase.On("FindByID", "1").Return(&usecase.User{
					ID:        "1",
					Email:     "test@example.com",
					FirstName: "Test",
					LastName:  "User",
				}, nil)
			},
		},
		{
			name:           "User not found",
			responseStatus: http.StatusNotFound,
			mock: func(mockUserUsecase *mocks.UserUsecase) {
				mockUserUsecase.On("FindByID", "1").Return(nil, nil)
			},
		},
		{
			name:           "Usecase Error",
			responseStatus: http.StatusInternalServerError,
			mock: func(mockUserUsecase *mocks.UserUsecase) {
				mockUserUsecase.On("FindByID", "1").Return(nil, errors.New("Usecase error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserUsecase := &mocks.UserUsecase{}
			service := &userService{mockUserUsecase}
			router := service.Routes()
			tt.mock(mockUserUsecase)

			req, _ := http.NewRequest("GET", "/1", nil)
			req.Header.Add("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req.WithContext(context.TODO()))

			if status := rr.Code; status != tt.responseStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.responseStatus)
			}
			mockUserUsecase.AssertExpectations(t)
		})
	}
}
