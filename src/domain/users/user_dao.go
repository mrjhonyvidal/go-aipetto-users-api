package users

import (
	"github.com/aipetto/go-aipetto-users-api/src/datasources/mysql/users_db"
	"github.com/aipetto/go-aipetto-users-api/src/utils/date_utils"
	"github.com/aipetto/go-aipetto-users-api/src/utils/errors"
	"github.com/aipetto/go-aipetto-users-api/src/utils/mysql_utils"
)

const (
	queryInsertUser 		= "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser 			= "SELECT id, first_name, last_name, email, date_created FROM users WHERE id =?;"
	queryUpdateUser			= "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser			= "DELETE FROM users WHERE id=?"
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
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		return mysql_utils.ParserError(getErr)
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
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if saveErr != nil {
		return mysql_utils.ParserError(saveErr)
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParserError(saveErr)
	}

	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysql_utils.ParserError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		return mysql_utils.ParserError(err)
	}
	return nil
}
