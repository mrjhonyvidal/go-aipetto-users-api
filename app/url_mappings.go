package app

import (
	"github.com/aipetto/go-aipetto-users-api/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}