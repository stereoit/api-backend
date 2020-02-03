package rest

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	mocks "github.com/stereoit/eventival/mocks/user/usecase"
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

			req, _ := http.NewRequest("POST", "/users/", bytes.NewBuffer([]byte(tt.requestPayload)))
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

