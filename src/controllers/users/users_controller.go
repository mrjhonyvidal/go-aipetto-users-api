package users

import (
	"net/http"
	"strconv"

	"github.com/aipetto/go-aipetto-users-api/src/domain/users"
	"github.com/aipetto/go-aipetto-users-api/src/services"
	"github.com/aipetto/go-aipetto-users-api/src/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	// Take the input request and validate it
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	// Call our service layer
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	// Call our service layer
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	// Get user_id from ULR Request
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	// User Id obtained from the URL
	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateUser(isPartial, user);
	if err != nil {
		restErr := errors.NewBadRequestError("")
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, result)
}