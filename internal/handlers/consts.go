package handlers

import "notes/pkg/logger"

var Log *logger.Loga = &logger.Log
var errors = map[int]string{
	400: "Validation error",
	404: "Not found",
	403: "Anauthorised",
	201: "Created",
	200: "Success",
	500: "Internal error",
}
