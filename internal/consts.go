package internal

import (
	"errors"
	"notes/pkg/logger"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var Log *logger.Loga = &logger.Log

var Errors = map[int]string{
	400: "Validation error",
	404: "Not found",
	403: "Unauthorised",
	201: "Created",
	200: "Success",
	500: "Internal error",
}

func LoggerParams(c *gin.Context, res_code int, t time.Time) []interface{} {
	params := []interface{}{c.Request.Method, c.Request.URL.Path, "-", res_code, Errors[res_code],
		strconv.FormatInt(time.Since(t).Milliseconds(), 10), "ms"}
	return params
}

// to compare ID from jwt and from path request
func CompareIDjwtPath(c *gin.Context, t time.Time) (int, error) {
	// get ID from JWT
	id_jwt, exists := c.Get("userID")
	if !exists {
		res_code := 500
		res_msg := "no userID in context"
		c.JSON(res_code, gin.H{
			"error": Errors[res_code],
		})
		Log.Error(res_msg, LoggerParams(c, res_code, t)...)
		return 0, errors.New(res_msg)
	}

	// get ID PATH
	i_id_path, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res_code := 404
		res_msg := "must be integer ID"
		c.JSON(res_code, gin.H{
			"error": Errors[res_code],
		})
		Log.Error(res_msg, LoggerParams(c, res_code, t)...)
		return 0, errors.New(res_msg)
	}

	// compare ID
	if id_jwt.(float64) != float64(i_id_path) {
		res_code := 403
		res_msg := "invalid ID"
		c.JSON(res_code, gin.H{
			"error": Errors[res_code],
		})
		Log.Error(res_msg, LoggerParams(c, res_code, t)...)
		return 0, errors.New(res_msg)
	}

	return i_id_path, nil
}
