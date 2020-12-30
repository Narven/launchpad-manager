package app

import (
	"fmt"
	"github.com/Narven/launchpad-manager/src/logger"
	"github.com/gin-gonic/gin"
	"os"
)

const (
	apiPort = "API_PORT"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("starting")

	router.Run(fmt.Sprintf(":%s", os.Getenv(apiPort)))
}
