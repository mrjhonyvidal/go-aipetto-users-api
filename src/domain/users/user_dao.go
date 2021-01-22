package users

import (
	"fmt"
	"github.com/aipetto/go-aipetto-users-api/src/datasources/mysql/users_db"
	"github.com/aipetto/go-aipetto-users-api/src/utils/date_utils"
	"github.com/aipetto/go-aipetto-users-api/src/utils/errors"
	"strings"
)

const (
	indexUniqueEmail	 	= "email_UNIQUE"
	errorNoRows				= "no rows in result set"
	queryInsertUser 		= "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser 			= "SELECT id, first_name, last_name, email, date_created FROM users WHERE id =?;"
)

func (user *User) Get() *errors.RestErr {

	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result  := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(
				fmt.Sprintf("user %d not found", user.Id))
		}
		fmt.Println(err)
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to get user %d:%s", user.Id, err.Error()))
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	user.DateCreated = date_utils.GetDateNowInStringFormat()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail){
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(
			fmt.Sprint("Error when trying to save user: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("Error when trying to save user: %s")
	}

	user.Id = userId
	return nil
}
