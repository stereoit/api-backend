package v1

import (
	"context"

	"github.com/stereoit/eventival/pkg/user/domain"
	"github.com/stereoit/eventival/pkg/user/interface/rpc/v1.0/protocol"
	"github.com/stereoit/eventival/pkg/user/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserService interface
type UserService interface {
	ListUser(ctx context.Context, in *protocol.ListUserRequestType) (*protocol.ListUserResponseType, error)
	RegisterUser(ctx context.Context, in *protocol.RegisterUserRequestType) (*protocol.RegisterUserResponseType, error)
	UpdateUser(ctx context.Context, in *protocol.UpdateUserRequestType) (*protocol.EmptyResponseType, error)
	DeleteUser(ctx context.Context, in *protocol.DeleteUserRequestType) (*protocol.EmptyResponseType, error)
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
	// TODO: parametrize GRPC service
	page := 0
	limit := 1
	users, err := s.userUsecase.ListAllUsers(page, limit)
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
		if _, ok := err.(*domain.DuplicateError); ok {
			err = status.Error(codes.AlreadyExists, err.Error())
		}
		return &protocol.RegisterUserResponseType{}, err
	}

	user, err := s.userUsecase.FindByID(newUserID)
	if err != nil {
		return &protocol.RegisterUserResponseType{}, err
	}

	user.FirstName = in.GetFirstName()
	user.LastName = in.GetLastName()

	if err := s.userUsecase.UpdateUser(user); err != nil {
		return &protocol.RegisterUserResponseType{}, err
	}

	return &protocol.RegisterUserResponseType{
		Id: newUserID,
	}, nil
}

func (s *userService) UpdateUser(ctx context.Context, in *protocol.UpdateUserRequestType) (*protocol.EmptyResponseType, error) {
	id := in.GetId()
	user, err := s.userUsecase.FindByID(id)
	if err != nil {
		return &protocol.EmptyResponseType{}, status.Error(codes.NotFound, err.Error())
	}
	user.FirstName = in.GetFirstName()
	user.LastName = in.GetLastName()

	if err := s.userUsecase.UpdateUser(user); err != nil {
		return &protocol.EmptyResponseType{}, status.Error(codes.Internal, err.Error())
	}

	return &protocol.EmptyResponseType{}, nil
}

func (s *userService) DeleteUser(ctx context.Context, in *protocol.DeleteUserRequestType) (*protocol.EmptyResponseType, error) {
	if err := s.userUsecase.DeleteUser(in.GetId()); err != nil {
		return &protocol.EmptyResponseType{}, err
	}
	return &protocol.EmptyResponseType{}, nil
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
