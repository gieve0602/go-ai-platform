package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	user := User{
		ID:   0o001,
		Name: "test1",
	}
	memRepo := NewMemRepository()
	ctx := context.Background()
	err := memRepo.Create(ctx, user)
	t.Logf("create user: %v\n", err)
	t.Logf("after create userData: %#v\n", memRepo.data)
}

func TestGetUser(t *testing.T) {
	memRepo := NewMemRepository()
	ctx := context.Background()
	tmp, err := memRepo.GetByID(ctx, 0o001)
	t.Logf("create user: %v, %v\n", tmp, err)
}

func TestGetNotExistsUser(t *testing.T) {
	memRepo := NewMemRepository()
	ctx := context.Background()
	user, err := memRepo.GetByID(ctx, 2)

	if err != nil {
		log.Printf("user not exists: %v\n", err)
	} else {
		log.Fatalf("user not exists err: %v, %v\n", user, err)
	}
}

func TestUpdateUser(t *testing.T) {
	user := User{
		ID:   0o001,
		Name: "test1",
	}
	memRepo := NewMemRepository()
	ctx := context.Background()
	err := memRepo.Create(ctx, user)
	if err != nil {
		t.Fatalf("Update user: get user err: %v\n", err)
	}
	user.Name = "test2"
	err = memRepo.Update(ctx, user)
	if err != nil {
		t.Fatalf("Update user error: %v\n", err)
	}
	_, err = memRepo.GetByID(ctx, 1)
	if err != nil {
		t.Fatalf("Update user: get user err: %v\n", err)
	} else {
		t.Logf("after Update userData: %#v\n", user)
	}
}

func TestDeleteUser(t *testing.T) {
	memRepo := NewMemRepository()
	ctx := context.Background()
	user := User{
		ID:   0o001,
		Name: "test1",
	}
	_ = memRepo.Create(ctx, user)
	err := memRepo.Delete(ctx, 1)
	if err != nil {
		t.Fatalf("Delete user error: %v\n", err)
	}
	t.Logf("after Delete userData: %#v\n", memRepo.data)
}

func TestRepeatCreate(t *testing.T) {
	user := User{
		ID:   0o001,
		Name: "test1",
	}
	memRepo := NewMemRepository()
	ctx := context.Background()
	err := memRepo.Create(ctx, user)
	if err != nil {
		t.Fatalf("create user fail: %v\n", err)
	}
	err = memRepo.Create(ctx, user)
	if err != nil {
		log.Printf("repeat create user fail: %v\n", err)
	}
	t.Logf("after repeat create userData: %#v\n", memRepo.data)
}

func TestParallelCreate(t *testing.T) {
	memRepo := NewMemRepository()
	ctx := context.Background()
	var wg sync.WaitGroup
	errChans := make(chan error, 10)
	var hasFail atomic.Bool
	for idx := range 10 {
		user := User{
			ID:   int64(idx),
			Name: fmt.Sprintf("test-%d", idx),
		}
		wg.Go(func() {
			err := memRepo.Create(ctx, user)
			if err != nil {
				hasFail.Store(true)
				errChans <- fmt.Errorf("parallel create user error: user-%d", idx)
			}
		})

	}
	wg.Wait()
	if hasFail.Load() {
		log.Fatal(<-errChans)
	}
	log.Printf("after parallel crate user: %v\n", memRepo.data)
}

func TestContextCancel(t *testing.T) {
	user := User{
		ID:   1,
		Name: "test1",
	}
	memRepo := NewMemRepository()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	cancel()
	err := memRepo.Create(ctx, user)
	if ctxErr := ctx.Err(); ctx.Err() != nil && !errors.Is(ctxErr, context.Canceled) {
		log.Fatalf("context cancel error: %v\n", ctxErr)
	}
	if err != nil {
		log.Printf("Get user error: %v\n", err)
	}
}

func BenchmarkMemoryRepositoryGetByID(b *testing.B) {
	memRepo := NewMemRepository()
	ctx := context.Background()
	for idx := range 10 {
		user := User{
			ID:   int64(idx),
			Name: fmt.Sprintf("test-%d", idx),
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

func BenchmarkMemoryRepositoryGetByID2(b *testing.B) {
	memRepo := NewMemRepository()
	ctx := context.Background()
	for idx := range 10 {
		user := User{
			ID:   int64(idx),
			Name: fmt.Sprintf("test-%d", idx),
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
