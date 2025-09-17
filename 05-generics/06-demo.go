package repo

import (
	"errors"
	"fmt"
)

type Repository[T any] struct {
	store map[int]T
}

func NewRepository[T any]() *Repository[T] {
	return &Repository[T]{store: make(map[int]T)}
}

// Error definitions
var (
	ErrNotFound     = errors.New("item not found")
	ErrAlreadyExist = errors.New("item already exists")
)

func (r *Repository[T]) Add(id int, value T) error {
	if _, exists := r.store[id]; exists {
		return ErrAlreadyExist
	}
	r.store[id] = value
	return nil
}

func (r *Repository[T]) Get(id int) (T, error) {
	val, ok := r.store[id]
	if !ok {
		var zero T
		return zero, ErrNotFound
	}
	return val, nil
}

func (r *Repository[T]) Delete(id int) error {
	if _, ok := r.store[id]; !ok {
		return ErrNotFound
	}
	delete(r.store, id)
	return nil
}

type User struct {
	ID   int
	Name string
}

func main() {
	userRepo := NewRepository[User]()

	err := userRepo.Add(1, User{ID: 1, Name: "Alice"})
	if err != nil {
		fmt.Println("Add Error:", err)
	}

	user, err := userRepo.Get(1)
	if err != nil {
		fmt.Println("Get Error:", err)
	} else {
		fmt.Println("Fetched:", user)
	}

	err = userRepo.Delete(2) // Attempt to delete non-existing user
	if err != nil {
		fmt.Println("Delete Error:", err)
	}
}
