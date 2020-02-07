package rest

import (
	"net/http"

	"github.com/sarulabs/di"

	"github.com/stereoit/eventival/pkg/user/domain"
	"github.com/stereoit/eventival/pkg/user/usecase"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Apply adds this resource to the given router
func Apply(path string, router chi.Router, ctn di.Container) {
	userService := New(ctn.Get("user-usecase").(usecase.UserUsecase))
	router.Mount(path, userService.Routes())
}

type userService struct {
	userUsecase usecase.UserUsecase
}

// UserService interfaces describes Application Service for Users
type UserService interface {
	Routes() chi.Router
}

// New returns instance of UserService
func New(userUsecase usecase.UserUsecase) UserService {
	return &userService{
		userUsecase: userUsecase,
	}
}

func (s *userService) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", s.listAllUsers)
	r.Post("/", s.registerUser)
	r.Patch("/{userID}", s.updateUser)
	r.Delete("/{userID}", s.deleteUser)

	return r
}

func (s *userService) listAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.userUsecase.ListAllUsers()
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	if err := render.RenderList(w, r, NewUsersListResponse(toUserList(users))); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (s *userService) registerUser(w http.ResponseWriter, r *http.Request) {
	userRequest := &RegisterUserRequest{}

	if err := render.Bind(r, userRequest); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, ErrBadRequest(err))
		return
	}

	newUserID, err := s.userUsecase.RegisterUser(userRequest.Email)
	if err != nil {
		if _, ok := err.(*domain.DuplicateError); ok {
			render.Render(w, r, ErrBadRequest(err))
			return
		}
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	user, err := s.userUsecase.FindByID(newUserID)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	// iterate over provided fields and try to set the updated one
	if userRequest.FirstName != nil {
		user.FirstName = *userRequest.FirstName
	}
	if userRequest.LastName != nil {
		user.LastName = *userRequest.LastName
	}

	if err := s.userUsecase.UpdateUser(user); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewRegisterUserResponse(newUserID))
}

func (s *userService) updateUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	user, err := s.userUsecase.FindByID(userID)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.Render(w, r, ErrNotFound)
		return
	}

	request := &UpdateRequest{}
	if err := render.Bind(r, request); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, ErrBadRequest(err))
		return
	}

	patchUser(user, request)

	if err := s.userUsecase.UpdateUser(user); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	render.Status(r, http.StatusNoContent)
	render.Render(w, r, NewEmptyResponse())
}

func patchUser(user *usecase.User, request *UpdateRequest) {
	// iterate over provided fields and try to set the updated one
	if request.FirstName != nil {
		user.FirstName = *request.FirstName
	}
	if request.LastName != nil {
		user.FirstName = *request.LastName
	}
}

func (s *userService) deleteUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	if err := s.userUsecase.DeleteUser(userID); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.Render(w, r, NewEmptyResponse())
		return
	}

	render.Status(r, http.StatusNoContent)
	render.Render(w, r, NewEmptyResponse())
}

type emptyResponse struct{}

// NewEmptyResponse returns empty response
func NewEmptyResponse() render.Renderer {
	return &emptyResponse{}
}
func (er *emptyResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
