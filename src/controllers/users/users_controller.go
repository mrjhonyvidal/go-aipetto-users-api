package users

import (
	"github.com/aipetto/go-aipetto-users-api/src/domain/users"
	"github.com/aipetto/go-aipetto-users-api/src/services"
	"github.com/aipetto/go-aipetto-users-api/src/utils/errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
)

func getUserIdFromUrl(userIdParam string) (int64, *errors.RestErr){
	// Convert the id from the url into a integer base 64
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}

func Create(c *gin.Context) {
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

func Get(c *gin.Context) {
	userId, idErr := getUserIdFromUrl(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
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

func Update(c *gin.Context) {
	userId, idErr := getUserIdFromUrl(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
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

func Delete(c *gin.Context) {
	userId, idErr := getUserIdFromUrl(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}