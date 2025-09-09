package handlers

import (
	"notes/inits"
	"notes/internal"
	"notes/internal/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetNote(c *gin.Context) {

	t := time.Now()

	// compare id
	i_ID, err := internal.CompareIDjwtPath(c, t)
	if err != nil {
		return
	}

	// get note id
	note_id := c.Param("note_id")
	i_note_id, err := strconv.Atoi(note_id)
	if err != nil {
		res_code := 400
		res_msg := "invalid note id"
		c.JSON(res_code, gin.H{
			"error": res_msg,
		})
		Log.Error(res_msg, internal.LoggerParams(c, res_code, t)...)
		return
	}

	// find note
	var note models.Note
	result := inits.DB.Where("id = ? and user_id = ?", i_note_id, i_ID).
		First(&note)

	if result.Error != nil {
		res_code := 500
		res_msg := "internal error"
		c.JSON(res_code, gin.H{
			"error": res_msg,
		})
		Log.Error(res_msg, internal.LoggerParams(c, res_code, t)...)
		return
	}

	// response with note
	res_code := 200
	c.JSON(res_code, gin.H{
		"note": note,
	})
	Log.Info("", internal.LoggerParams(c, res_code, t)...)

}
