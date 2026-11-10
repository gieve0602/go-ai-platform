package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassNilRepository(t *testing.T) {
	_, err := NewUserService(nil)
	assert.ErrorIs(t, err, ErrValidRepo)
}

func TestServiceCreateUser(t *testing.T) {
	repo := NewMemRepository()
	svc, err := NewUserService(repo)
	assert.NoError(t, err)
	ctx := context.Background()
	user := User{
		ID:   int64(0),
		Name: "",
	}
	err = svc.Create(ctx, user)
	assert.ErrorIs(t, err, ErrValidParam)
	user.ID = -2
	err = svc.Create(ctx, user)
	assert.ErrorIs(t, err, ErrValidParam)
	user.ID = 2
	err = svc.Create(ctx, user)
	assert.ErrorIs(t, err, ErrValidParam)
	user.Name = "test-2"
	err = svc.Create(ctx, user)
	assert.ErrorIs(t, err, ErrValidParam)
	user.Password = "123456"
	err = svc.Create(ctx, user)
	assert.NoError(t, err)
}

func TestServiceUpdateUser(t *testing.T) {
	repo := NewMemRepository()
	svc, err := NewUserService(repo)
	assert.NoError(t, err)

	ctx := context.Background()
	user := User{
		ID:       int64(1),
		Name:     "test-1",
		Password: "123456",
	}
	err = svc.Create(ctx, user)
	assert.NoError(t, err)
	user.ID = 1
	user.Name = "test-2"
	user.Password = "1234567"
	err = svc.Update(ctx, user)
	assert.NoError(t, err)
	newUser, err := svc.GetByID(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, newUser, user)
}

func TestServiceDeleteUser(t *testing.T) {
	repo := NewMemRepository()
	svc, err := NewUserService(repo)
	assert.NoError(t, err)

	ctx := context.Background()
	user := User{
		ID:       int64(1),
		Name:     "test-1",
		Password: "123456",
	}
	err = svc.Create(ctx, user)
	assert.NoError(t, err)
	err = svc.Delete(ctx, user.ID)
	assert.NoError(t, err)
	_, err = svc.GetByID(ctx, user.ID)
	assert.ErrorIs(t, err, ErrUserNotExists)
}

func TestServiceGetUser(t *testing.T) {
	repo := NewMemRepository()
	svc, err := NewUserService(repo)
	assert.NoError(t, err)

	ctx := context.Background()
	user := User{
		ID:       int64(1),
		Name:     "test-1",
		Password: "123456",
	}
	err = svc.Create(ctx, user)
	assert.NoError(t, err)
	newUser, err := svc.GetByID(ctx, user.ID)
	assert.NoError(t, err)
	assert.Equal(t, newUser, user)
}

func TestServiceReturnRepoError(t *testing.T) {
	repo := NewMemRepository()
	svc, err := NewUserService(repo)
	assert.NoError(t, err)

	ctx := context.Background()
	user := User{
		ID:       int64(1),
		Name:     "test-1",
		Password: "12345",
	}
	err = svc.Create(ctx, user)
	assert.ErrorIs(t, err, ErrWeakPassword)
}
