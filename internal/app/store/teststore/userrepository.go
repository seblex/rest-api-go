package teststore

import (
	"github.com/seblex/rest-api-go/internal/app/model"
	"github.com/seblex/rest-api-go/internal/app/store"
)

// TestUserRepository ...
type UserRepository struct {
	store *Store
	users map[string]*model.User
}

// Create ...
func (r *UserRepository) Create(u *model.User) error {
	err := u.Validate()
	if err != nil {
		return err
	}

	err = u.BeforeCreate()
	if err != nil {
		return err
	}

	r.users[u.Email] = u
	u.ID = len(r.users)

	return nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u, ok := r.users[email]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}