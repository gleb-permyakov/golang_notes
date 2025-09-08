package internal

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

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
