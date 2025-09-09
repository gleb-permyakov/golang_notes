package handlers

import (
	"notes/inits"
	"notes/internal"
	"notes/internal/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func DeleteNote(c *gin.Context) {
	t := time.Now()

	// compare ID
	_, err := internal.CompareIDjwtPath(c, t)
	if err != nil {
		return
	}

	// get note_id
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

	// delete the note
	var note models.Note
	result := inits.DB.Delete(&note, i_note_id)
	if result.Error != nil {
		res_code := 500
		res_msg := "internal error"
		c.JSON(res_code, gin.H{
			"error": res_msg,
		})
		Log.Error("can not delete the note", internal.LoggerParams(c, res_code, t)...)
		return
	}

	res_code := 200
	res_msg := "note deleted"
	c.JSON(res_code, gin.H{
		"message": res_msg,
		"note_id": result.RowsAffected,
	})
	Log.Info("", internal.LoggerParams(c, res_code, t)...)

}
