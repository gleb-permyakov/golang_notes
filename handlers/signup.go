package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {

	var user struct {
		id         int
		username   string
		created_at string
	}

	// read body
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
	}

	// add to db

	// add jwt

	//

}
