package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGobitly(c *gin.Context) {
	// datastore.Connect()

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})

}