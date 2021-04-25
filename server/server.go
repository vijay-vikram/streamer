package server

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// Serve loads the application's config file and starts the server
func Serve() {
	mapUrls()
	router.Run(":8081")
}
