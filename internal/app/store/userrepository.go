package store

import "github.com/seblex/rest-api-go/internal/app/model"

// User repository ...
type UserRepository struct {
	store *Store
}

// Create user ...
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	err := u.Validate()
	if err != nil {
		return nil, err
	}

	err = u.BeforeCreate()
	if err != nil {
		return nil, err
	}

	err = r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// Find user by email ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	)

	if err != nil {
		return nil, err
	}

	return u, nil
}