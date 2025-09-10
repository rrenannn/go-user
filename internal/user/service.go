package user

import (
	"context"
	"fmt"

	db "github.com/rrenannn/go-user/db/sqlc"
	"github.com/rrenannn/go-user/infra/crypt"
)

type ServiceInterface interface {
	CreateUser(ctx context.Context, data UserRequest) (UserResponse, error)
	GetUserById(ctx context.Context, id int64) (UserResponse, error)
	GetUserByEmail(ctx context.Context, email string) (UserResponse, error)
	ResetPassword(ctx context.Context, arg db.ResetPasswordParams) error
}

type Service struct {
	repo  Repository
	crypt *crypt.Crypt
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(ctx context.Context, data UserRequest) (UserResponse, error) {
	var response UserResponse

	existingUser, err := s.repo.GetUserByEmail(ctx, data.Email)
	if err == nil && existingUser.ID != 0 {
		return UserResponse{}, fmt.Errorf("user with email %s already exists", data.Email)
	}

	hashedPassword, err := s.crypt.HashPassword(data.Password)
	if err != nil {
		return UserResponse{}, err
	}

	user, err := s.repo.CreateUser(ctx, db.CreateUserParams{
		Name:     data.Name,
		Email:    data.Email,
		Password: hashedPassword,
		Status:   data.Status,
	})
	if err != nil {
		return UserResponse{}, err
	}

	response = UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
	}

	return response, nil
}

func (s *Service) GetUserById(ctx context.Context, id int64) (UserResponse, error) {
	user, err := s.repo.GetUserById(ctx, id)
	if err != nil {
		return UserResponse{}, err
	}

	response := UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
	}

	return response, nil
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (UserResponse, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return UserResponse{}, err
	}
	response := UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: &user.UpdatedAt.Time,
	}

	return response, nil
}

func (s *Service) ResetPassword(ctx context.Context, arg db.ResetPasswordParams) error {
	err := s.repo.ResetPassword(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}
