package users

import (
	"errors"
	"fmt"
	"github.com/aipetto/go-aipetto-users-api/src/datasources/mysql/users_db"
	"github.com/aipetto/go-aipetto-users-api/src/logger"
	"github.com/aipetto/go-aipetto-users-api/src/utils/crypto_utils"
	"github.com/aipetto/go-aipetto-utils/src/rest_errors"
	"github.com/aipetto/go-aipetto-users-api/src/utils/mysql_utils"
	"strings"
)

const (
	queryInsertUser 				= "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser 					= "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id =?;"
	queryUpdateUser					= "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser					= "DELETE FROM users WHERE id=?"
	queryFindUserByStatus			= "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindByEmailAndStatusActive = "SELECT id, first_name, last_name, email, date_created, status, password FROM users WHERE email=? AND status=?;"
)

func (user *User) Get() *rest_errors.RestErr {

	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return rest_errors.NewInternalServerError("database error", errors.New("database error"))
	}
	defer stmt.Close()

	result  := stmt.QueryRow(user.Id)
	if getErr := result.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.DateCreated,
		&user.Status); err != nil {
		logger.Error("error when trying to get user by id", getErr)
		return rest_errors.NewInternalServerError("database error", errors.New("database error"))
	}
	return nil
}

func (user *User) FindByEmailAndPassword(requestPassword string) *rest_errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryFindByEmailAndStatusActive)
	if err != nil {
		logger.Error("error when trying to prepare get user by email and password", err)
		return rest_errors.NewInternalServerError("database error", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, StatusActive)
	if getErr := result.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.DateCreated,
		&user.Status,
		&user.Password); getErr != nil {
		if strings.Contains(getErr.Error(), mysql_utils.ErrorNoRows) {
			return rest_errors.InvalidCredentialsError("invalid user credentials")
		}
	}

	if verifyPasswordErr := crypto_utils.VerifyPassword(user.Password, requestPassword); verifyPasswordErr != nil {
			return rest_errors.NewInternalServerError("invalid password", errors.New("server authentication error"))
	}
	return nil
}

func (user *User) Save() *rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare the sql save statement", err)
		return rest_errors.NewInternalServerError("database error", errors.New("database error"))
	}

	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(
		user.FirstName,
		user.LastName,
		user.Email,
		user.DateCreated,
		user.Status,
		user.Password)

	if saveErr != nil {
		logger.Error("error when trying to save the user", saveErr)
		return rest_errors.NewInternalServerError("database error", errors.New("database error"))
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last user id", err)
		return rest_errors.NewInternalServerError("database error", errors.New("database_error"))
	}

	user.Id = userId
	return nil
}

func (user *User) Update() *rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update statement", err)
		return rest_errors.NewInternalServerError("database error", errors.New("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("error when trying to sql update", err)
		return rest_errors.NewInternalServerError("database error", errors.New("database error"))
	}
	return nil
}

func (user *User) Delete() *rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return rest_errors.NewInternalServerError("database error", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		logger.Error("error when trying to delete the user", err)
		return rest_errors.NewInternalServerError("database error", errors.New("database errors"))
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *rest_errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find the user by id statement", err)
		return nil, rest_errors.NewInternalServerError("database error", errors.New("database error"))
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find user by id", err)
		return nil, rest_errors.NewInternalServerError("database error", errors.New("database error"))
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next(){
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when scan user row into user struct", err)
			return nil, rest_errors.NewInternalServerError("database error", errors.New("database error"))
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no users mathcing status %s", status))
	}

	return results, nil
}
