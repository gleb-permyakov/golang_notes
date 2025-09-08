package handlers

import (
	"notes/inits"
	"notes/internal"
	"notes/internal/models"
	"strconv"
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

	// get ID from JWT
	id_jwt, exists := c.Get("userID")
	if !exists {
		res_code := 500
		res_msg := "no userID in context"
		c.JSON(res_code, gin.H{
			"error": internal.Errors[res_code],
		})
		Log.Error(res_msg, internal.LoggerParams(c, res_code, t)...)
		return
	}

	// get ID PATH
	i_id_path, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res_code := 404
		res_msg := "must be integer ID"
		c.JSON(res_code, gin.H{
			"error": internal.Errors[res_code],
		})
		Log.Error(res_msg, internal.LoggerParams(c, res_code, t)...)
		return
	}

	// compare ID
	if id_jwt.(float64) != float64(i_id_path) {
		res_code := 403
		res_msg := "invalid ID"
		c.JSON(res_code, gin.H{
			"error": internal.Errors[res_code],
		})
		Log.Error(res_msg, internal.LoggerParams(c, res_code, t)...)
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
		res_code := 200
		res_msg := "created note"
		c.JSON(res_code, gin.H{
			"error": res_msg,
		})
		Log.Info(res_msg, internal.LoggerParams(c, res_code, t)...)
		return
	}

}
