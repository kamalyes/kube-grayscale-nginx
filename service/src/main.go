package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var version = flag.String("v", "blue", "blue")

func main() {
	router := gin.Default()
	router.GET("", func(c *gin.Context) {
		flag.Parse()
		hostname, _ := os.Hostname()
		c.String(http.StatusOK, "This is version:%s running in pod %s", *version, hostname)
	})
	router.Run(":8080")
}
