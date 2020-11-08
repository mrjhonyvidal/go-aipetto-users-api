package users

import (
	"encoding/json"
	"fmt"
	"github.com/aipetto/go-aipetto-users-api/domain/users"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// TODO: Handler error
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err.Error())
		// TODO Handle json error
		return
	}
	fmt.Println(user)
	fmt.Println(string(bytes))
	fmt.Println(err)
	c.String(http.StatusNotImplemented, "to implement")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "to implement")
}