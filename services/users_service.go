package services

import (
	"github.com/aipetto/go-aipetto-users-api/domain/users"
)

func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
