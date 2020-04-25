package store

import (
	"github.com/seblex/rest-api-go/internal/app/model"
)

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}