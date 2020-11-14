package services

import (
	"net/http"

	"github.com/aipetto/go-aipetto-users-api/domain/users"
	"github.com/aipetto/go-aipetto-users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	return &user, nil

	return &user, &errors.RestErr{
		Status: http.StatusInternalServerError,
	}
}
