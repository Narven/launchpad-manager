package app

import (
	"fmt"
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

	router.Run(fmt.Sprintf(":%s", os.Getenv(apiPort)))
}
