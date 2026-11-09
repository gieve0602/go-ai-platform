package user

import (
	"context"
	"errors"
	"sync"
)

var (
	ErrValidUser     = errors.New("user is nil")
	ErrUserExists    = errors.New("user exists")
	ErrUserNotExists = errors.New("user not exists")
	ErrValidParam    = errors.New("invalid parameter")
)

type (
	MemRepository struct {
		data map[int64]User
		mux  sync.RWMutex
	}
)

func NewMemRepository() *MemRepository {
	return &MemRepository{
		data: make(map[int64]User),
	}
}

func (m *MemRepository) Create(ctx context.Context, user User) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if user == (User{}) {
		return ErrValidUser
	}
	m.mux.Lock()
	defer m.mux.Unlock()
	if _, ok := m.data[user.ID]; ok {
		return ErrUserExists
	}
	m.data[user.ID] = user
	return nil
}

func (m *MemRepository) Update(ctx context.Context, user User) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if user == (User{}) {
		return ErrValidUser
	}
	m.mux.Lock()
	defer m.mux.Unlock()
	if _, ok := m.data[user.ID]; !ok {
		return ErrUserNotExists
	}
	m.data[user.ID] = user
	return nil
}

func (m *MemRepository) Delete(ctx context.Context, id int64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if id < 0 {
		return ErrValidParam
	}
	m.mux.Lock()
	defer m.mux.Unlock()
	if _, ok := m.data[id]; !ok {
		return ErrUserNotExists
	}
	delete(m.data, id)
	return nil
}

func (m *MemRepository) GetByID(ctx context.Context, id int64) (User, error) {
	if err := ctx.Err(); err != nil {
		return User{}, err
	}
	if id < 0 {
		return User{}, ErrValidParam
	}
	m.mux.RLock()
	defer m.mux.RUnlock()
	if user, ok := m.data[id]; ok {
		return user, nil
	} else {
		return User{}, ErrUserNotExists
	}
}

func (m *MemRepository) GetByID2(ctx context.Context, id int64) (User, error) {
	if err := ctx.Err(); err != nil {
		return User{}, err
	}
	if id < 0 {
		return User{}, ErrValidParam
	}
	m.mux.Lock()
	defer m.mux.Unlock()
	if user, ok := m.data[id]; ok {
		return user, nil
	} else {
		return User{}, ErrUserNotExists
	}
}
