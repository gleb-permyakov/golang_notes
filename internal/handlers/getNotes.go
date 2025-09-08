package handlers

import (
	"notes/inits"
	"notes/internal"
	"notes/internal/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetNotes(c *gin.Context) {

	t := time.Now()

	// compare id
	i_ID, err := internal.CompareIDjwtPath(c, t)
	if err != nil {
		return
	}

	// params of request
	s_limit := c.DefaultQuery("limit", "10")
	s_offset := c.DefaultQuery("offset", "0")
	s_sort := c.DefaultQuery("sort", "asc")

	// assertions
	i_limit, err := strconv.Atoi(s_limit)
	if err != nil || i_limit <= 0 {
		i_limit = 10
	}

	i_offset, err := strconv.Atoi(s_offset)
	if err != nil || i_offset < 0 {
		i_offset = 0
	}

	if s_sort != "desc" {
		s_sort = "ASC"
	} else {
		s_sort = "DESC"
	}

	// get notes
	var all_user_notes []models.Note
	result := inits.DB.Where("user_id = ?", i_ID).
		Limit(i_limit).
		Offset(i_offset).
		Order("created_at " + s_sort).
		Find(&all_user_notes)

	if result.Error != nil {
		res_code := 500
		res_msg := result.Error.Error()
		c.JSON(res_code, gin.H{
			"error": "internal error",
		})
		Log.Error(res_msg, internal.LoggerParams(c, res_code, t)...)
		return
	}

	// response with all notes
	// thats how to count all notes
	// var num_notes int64
	// inits.DB.Model(&models.Note{}).Where("user_id = ?", i_ID).Count(&num_notes)
	res_code := 200
	c.JSON(res_code, gin.H{
		"notes": all_user_notes,
	})
	Log.Info("", internal.LoggerParams(c, res_code, t)...)

}
