package v1

import (
	"context"

	"github.com/stereoit/eventival/pkg/user/interface/rpc/v1.0/protocol"
	"github.com/stereoit/eventival/pkg/user/usecase"
)

// UserService interface
type UserService interface {
	ListUser(ctx context.Context, in *protocol.ListUserRequestType) (*protocol.ListUserResponseType, error)
	RegisterUser(ctx context.Context, in *protocol.RegisterUserRequestType) (*protocol.RegisterUserResponseType, error)
}

type userService struct {
	userUsecase usecase.UserUsecase
}

// NewUserService returns new instance of implementation of UserService protocol
func NewUserService(userUsecase usecase.UserUsecase) UserService {
	return &userService{
		userUsecase: userUsecase,
	}
}

func (s *userService) ListUser(ctx context.Context, in *protocol.ListUserRequestType) (*protocol.ListUserResponseType, error) {
	users, err := s.userUsecase.ListAllUsers()
	if err != nil {
		return nil, err
	}

	res := &protocol.ListUserResponseType{
		Users: toProtocolUserList(users),
	}

	return res, nil

}

func (s *userService) RegisterUser(ctx context.Context, in *protocol.RegisterUserRequestType) (*protocol.RegisterUserResponseType, error) {

	newUserID, err := s.userUsecase.RegisterUser(in.GetEmail())
	if err != nil {
		return &protocol.RegisterUserResponseType{}, err
	}

	return &protocol.RegisterUserResponseType{
		Id: newUserID,
	}, nil
}

func toProtocolUser(user *usecase.User) *protocol.User {
	return &protocol.User{
		Id:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func toProtocolUserList(users []*usecase.User) []*protocol.User {
	res := make([]*protocol.User, len(users))
	for i, user := range users {
		res[i] = toProtocolUser(user)
	}
	return res
}
