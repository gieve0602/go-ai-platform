package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassNilRepository(t *testing.T) {
	svc := NewUserService(nil)
	assert.Nil(t, svc, nil)
}

func TestServiceCreateUser(t *testing.T) {
	repo := NewMemRepository()
	svc := NewUserService(repo)

	ctx := context.Background()
	user := User{
		ID:   int64(-2),
		Name: "",
	}
	err := svc.Create(ctx, user)
	assert.ErrorIs(t, err, ErrValidParam)
}

func TestServiceUpdateUser(t *testing.T) {
	repo := NewMemRepository()
	svc := NewUserService(repo)

	ctx := context.Background()
	user := User{
		ID:       int64(1),
		Name:     "test-1",
		Password: "12345",
	}
	err := svc.Create(ctx, user)
	assert.ErrorIs(t, err, nil)
	user.ID = 1
	user.Name = "test-2"
	user.Password = "123456"
	err = svc.Update(ctx, user)
	assert.ErrorIs(t, err, nil)
	newUser, err := svc.GetByID(ctx, 1)
	assert.ErrorIs(t, err, nil)
	assert.Equal(t, newUser, user)
}

func TestServiceDeleteUser(t *testing.T) {
	repo := NewMemRepository()
	svc := NewUserService(repo)

	ctx := context.Background()
	user := User{
		ID:       int64(1),
		Name:     "test-1",
		Password: "12345",
	}
	err := svc.Create(ctx, user)
	assert.ErrorIs(t, err, nil)
	err = svc.Delete(ctx, user.ID)
	assert.ErrorIs(t, err, nil)
	_, err = svc.GetByID(ctx, user.ID)
	assert.ErrorIs(t, err, ErrUserNotExists)
}
func TestServiceGetUser(t *testing.T) {}
