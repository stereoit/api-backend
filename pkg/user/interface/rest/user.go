package rest

import (
	"errors"
	"net/http"

	"github.com/stereoit/eventival/pkg/user/usecase"

	"github.com/go-chi/render"
)

// User represents single user
type User struct {
	ID        string `json:"id,omitempty"`
	Email     string `json:"email"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

// RegisterUserRequest represent incoming user object
type RegisterUserRequest struct {
	*User
}

// Bind implements Binder interface
func (u *RegisterUserRequest) Bind(r *http.Request) error {
	if u.User == nil {
		return errors.New("missing required fields")
	}

	if u.Email == "" {
		return errors.New("missing required User field")
	}

	return nil
}

// UpdateRequest maps complete user
type UpdateRequest struct {
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
}

// Bind implements render interface
func (u *UpdateRequest) Bind(r *http.Request) error {
	return nil
}

// UserResponse represents response to created user resource
type UserResponse struct {
	*User
}

// Render implements Renderer interface, not sure if needed for empty implementation
func (u *UserResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// NewUserResponse returns I do not know
func NewUserResponse(user *User) *UserResponse {
	resp := &UserResponse{
		user,
	}

	// if resp.User == nil {
	// 	if user, _ := dbGetUser(resp.UserID); user != nil {
	// 		resp.User = NewUserPayloadResponse(user)
	// 	}
	// }

	return resp
}

// RegisterUserResponse returns just an ID of the new user
type RegisterUserResponse struct {
	ID string `json:"id"`
}

// NewRegisterUserResponse returns ID
func NewRegisterUserResponse(id string) *RegisterUserResponse {
	return &RegisterUserResponse{
		ID: id,
	}
}

// Render implements Renderer interface, not sure if needed for empty implementation
func (u *RegisterUserResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// NewUsersListResponse something
func NewUsersListResponse(users []*User) []render.Renderer {
	list := []render.Renderer{}
	for _, user := range users {
		list = append(list, NewUserResponse(user))
	}
	return list
}

// converts between entities
func toUser(user *usecase.User) *User {
	return &User{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

// converts between entities
func toUserList(users []*usecase.User) []*User {
	res := make([]*User, len(users))
	for i, user := range users {
		res[i] = toUser(user)
	}
	return res
}
