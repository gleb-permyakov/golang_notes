package handlers

import (
	"github.com/gin-gonic/gin"
)

func CreateNote(c *gin.Context) {

	var body struct {
		Title   string
		Content string
	}

	err := c.Bind(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "failed to red body",
		})
	}

}
