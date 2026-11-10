package user

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createExampleData(t *testing.T, ctx context.Context, memRepo *MemRepository) {
	user1 := User{
		ID:       1,
		Name:     "test1",
		Password: "123456",
	}
	user2 := User{
		ID:       2,
		Name:     "test2",
		Password: "123456",
	}
	user3 := User{
		ID:       3,
		Name:     "test3",
		Password: "123456",
	}
	err := memRepo.Create(ctx, user1)
	assert.Equal(t, err, nil)
	err = memRepo.Create(ctx, user2)
	assert.Equal(t, err, nil)
	err = memRepo.Create(ctx, user3)
	assert.Equal(t, err, nil)
}

func TestCreateUser(t *testing.T) {
	user := User{
		ID:       0o001,
		Name:     "test1",
		Password: "123456",
	}
	memRepo := NewMemRepository()
	ctx := context.Background()
	err := memRepo.Create(ctx, user)
	assert.NoError(t, err)
	tmpUser, err := memRepo.GetByID(ctx, user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user, tmpUser)
}

func TestGetUser(t *testing.T) {
	memRepo := NewMemRepository()
	ctx := context.Background()
	createExampleData(t, ctx, memRepo)
	tmp, err := memRepo.GetByID(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, tmp.ID, int64(1))
}

func TestGetNotExistsUser(t *testing.T) {
	memRepo := NewMemRepository()
	ctx := context.Background()
	_, err := memRepo.GetByID(ctx, 2)
	assert.ErrorIs(t, err, ErrUserNotExists)
}

func TestUpdateUser(t *testing.T) {
	memRepo := NewMemRepository()
	ctx := context.Background()
	createExampleData(t, ctx, memRepo)
	user, err := memRepo.GetByID(ctx, 1)
	assert.NoError(t, err)
	user.Name = "test1-1"
	err = memRepo.Update(ctx, user)
	assert.NoError(t, err)
	newUser, err := memRepo.GetByID(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, newUser.Name, user.Name)
}

func TestDeleteUser(t *testing.T) {
	memRepo := NewMemRepository()
	ctx := context.Background()
	createExampleData(t, ctx, memRepo)
	err := memRepo.Delete(ctx, 1)
	assert.NoError(t, err)
	_, err = memRepo.GetByID(ctx, 1)
	assert.ErrorIs(t, err, ErrUserNotExists)
}

func TestRepeatCreate(t *testing.T) {
	memRepo := NewMemRepository()
	ctx := context.Background()
	createExampleData(t, ctx, memRepo)
	user1 := User{
		ID:       1,
		Name:     "test1",
		Password: "123456",
	}
	err := memRepo.Create(ctx, user1)
	assert.ErrorIs(t, err, ErrUserExists)
}

func TestParallelCreate(t *testing.T) {
	memRepo := NewMemRepository()
	ctx := context.Background()
	count := 10
	var wg sync.WaitGroup
	for idx := range count {
		idx++
		user := User{
			ID:       int64(idx),
			Name:     fmt.Sprintf("test-%d", idx),
			Password: "123456",
		}
		wg.Go(func() {
			err := memRepo.Create(ctx, user)
			assert.NoError(t, err)
		})
	}
	wg.Wait()
	assert.Equal(t, len(memRepo.data), count)
}

func TestContextCancel(t *testing.T) {
	user := User{
		ID:       1,
		Name:     "test1",
		Password: "123456",
	}
	memRepo := NewMemRepository()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := memRepo.Create(ctx, user)
	assert.ErrorIs(t, err, context.Canceled)
}

func BenchmarkMemoryRepositoryGetByID(b *testing.B) {
	memRepo := NewMemRepository()
	ctx := context.Background()
	for idx := range 10 {
		idx++
		user := User{
			ID:       int64(idx),
			Name:     fmt.Sprintf("test-%d", idx),
			Password: "123456",
		}
		_ = memRepo.Create(ctx, user)
	}
	b.ResetTimer()
	for b.Loop() {
		for idx := range 10 {
			_, _ = memRepo.GetByID(ctx, int64(idx))
		}
	}
}
