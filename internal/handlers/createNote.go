package handlers

import (
	"notes/inits"
	"notes/internal"
	"notes/internal/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateNote(c *gin.Context) {

	t := time.Now()

	var body struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	// read body
	err := c.Bind(&body)
	if err != nil {
		res_code := 400
		res_msg := "failed to read body"
		c.JSON(res_code, gin.H{
			"error": res_msg,
		})
		Log.Error(res_msg, internal.LoggerParams(c, res_code, t)...)
		return
	}

	// compare ID
	i_id_path, err := internal.CompareIDjwtPath(c, t)
	if err != nil {
		return
	}

	// make note
	note := models.Note{User_id: i_id_path, Title: body.Title, Content: body.Content}
	result := inits.DB.Create(&note)

	if result.Error != nil {
		res_code := 403
		res_msg := result.Error.Error()
		c.JSON(res_code, gin.H{
			"error": "can not create note",
		})
		Log.Error(res_msg+" cannot create the note", internal.LoggerParams(c, res_code, t)...)
		return
	} else {
		res_code := 201
		res_msg := "created note"
		c.JSON(res_code, gin.H{
			"message": res_msg,
		})
		Log.Info(res_msg, internal.LoggerParams(c, res_code, t)...)
		return
	}

}
