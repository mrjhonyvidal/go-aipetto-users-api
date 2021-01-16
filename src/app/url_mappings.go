package app

import (
	"github.com/aipetto/go-aipetto-users-api/src/controllers/ping"
	"github.com/aipetto/go-aipetto-users-api/src/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}