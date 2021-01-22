package mysql_utils

import (
	"fmt"
	"github.com/aipetto/go-aipetto-users-api/src/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	errorNoRows		= "no rows in result set"
)

func ParserError(err error) *errors.RestErr {
	sqlErr, convertedErrorToMySQLError := err.(*mysql.MySQLError)

	if !convertedErrorToMySQLError {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError(fmt.Sprintf("invalid data"))
	}
	return errors.NewInternalServerError("error processing request")
}
