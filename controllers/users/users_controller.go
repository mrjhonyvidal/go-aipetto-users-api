package users

import (
	"fmt"
	"net/http"

	"github.com/aipetto/go-aipetto-users-api/domain/users"
	"github.com/aipetto/go-aipetto-users-api/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		// TODO: return bad request to the caller
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// TODO Handle user creation
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "to implement")
}
