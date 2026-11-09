package user

import "context"

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(ctx context.Context, user User) error {
	if user == (User{}) {
		return ErrValidUser
	}
	return s.repo.Create(ctx, user)
}

func (s *UserService) Update(ctx context.Context, user User) error {
	if user == (User{}) {
		return ErrValidUser
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
