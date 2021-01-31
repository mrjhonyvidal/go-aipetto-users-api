package mysql_utils

import (
	"errors"
	"fmt"
	"github.com/aipetto/go-aipetto-utils/src/rest_errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	ErrorNoRows		= "no rows in result set"
)

func ParserError(err error) *rest_errors.RestErr {
	sqlErr, convertedErrorToMySQLError := err.(*mysql.MySQLError)

	if !convertedErrorToMySQLError {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return rest_errors.NewNotFoundError("no record matching given id")
		}
		return rest_errors.NewInternalServerError("error parsing database response", errors.New("database parsing error"))
	}

	switch sqlErr.Number {
	case 1062:
		return rest_errors.NewBadRequestError(fmt.Sprintf("invalid data"))
	}
	return rest_errors.NewInternalServerError("error processing request", errors.New("request process error"))
}
