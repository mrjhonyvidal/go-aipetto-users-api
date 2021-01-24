package app

import (
	"github.com/aipetto/go-aipetto-users-api/src/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("Starting AIPETTO Application...")
	router.Run(":8081")
}