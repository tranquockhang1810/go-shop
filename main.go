package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Lab03 - tele",
		})
	})
	r.Run(":3000")
}

// docker exec -it --user root jenkins-go /bin/bash
// chmod 666 /var/run/docker.sock
