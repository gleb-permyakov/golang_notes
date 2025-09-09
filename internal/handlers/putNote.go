package handlers

import (
	"notes/inits"
	"notes/internal"
	"notes/internal/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func PutNote(c *gin.Context) {
	t := time.Now()

	// compare ID
	_, err := internal.CompareIDjwtPath(c, t)
	if err != nil {
		return
	}

	// bind the body
	var body struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	err = c.Bind(&body)
	if err != nil {
		res_code := 400
		res_msg := "body is invalid"
		c.JSON(res_code, gin.H{
			"error": res_msg,
		})
		Log.Error(res_msg, internal.LoggerParams(c, res_code, t)...)
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

	// find the note
	var note models.Note
	inits.DB.Where("id = ?", i_note_id).
		First(&note)

	// update note data
	if body.Content != "" {
		note.Content = body.Content
	}
	if body.Title != "" {
		note.Title = body.Title
	}

	note.UpdatedAt = time.Now()

	// put the note
	result := inits.DB.Save(&note)
	if result.Error != nil {
		res_code := 500
		res_msg := "internal error"
		c.JSON(res_code, gin.H{
			"error": res_msg,
		})
		Log.Error("can not put the note", internal.LoggerParams(c, res_code, t)...)
		return
	}

	res_code := 200
	res_msg := "note updated"
	c.JSON(res_code, gin.H{
		"message":     res_msg,
		"new_title":   note.Title,
		"new_content": note.Content,
		"updated_at":  note.UpdatedAt,
	})
	Log.Info("", internal.LoggerParams(c, res_code, t)...)

}
