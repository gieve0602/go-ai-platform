package user

import (
	"context"
	"errors"
	"strings"
)

var (
	ErrValidPassword = errors.New("password is not valid")
	ErrValidRepo     = errors.New("Repository is nil")
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	if repo == nil {
		return nil
	}
	return &UserService{
		repo: repo,
	}
}

func validateUser(user User) error {
	if user.ID < 0 {
		return ErrValidParam
	}
	if strings.TrimSpace(user.Name) == "" {
		return ErrValidUser
	}
	if strings.TrimSpace(user.Password) == "" {
		return ErrValidPassword
	}
	return nil
}

func (s *UserService) Create(ctx context.Context, user User) error {
	if err := validateUser(user); err != nil {
		return err
	}
	return s.repo.Create(ctx, user)
}

func (s *UserService) Update(ctx context.Context, user User) error {
	if err := validateUser(user); err != nil {
		return err
	}
	return s.repo.Update(ctx, user)
}

func (s *UserService) Delete(ctx context.Context, id int64) error {
	if id < 0 {
		return ErrValidParam
	}
	return s.repo.Delete(ctx, id)
}

func (s *UserService) GetByID(ctx context.Context, id int64) (User, error) {
	if id < 0 {
		return User{}, ErrValidParam
	}
	return s.repo.GetByID(ctx, id)
}
