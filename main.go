package main

import (
	"io"
	"notes/inits"
	"notes/internal/handlers"
	"notes/pkg/logger"
	"os"

	"github.com/gin-gonic/gin"
)

var Log *logger.Loga = &logger.Log

func init() {
	inits.EnvInit()
	logger.New()
	inits.ConnectDB()
	inits.DoTablesDB()
}

func main() {

	// logger
	// Log := logger.Log
	Log.Warn("DEV Level")

	// Полностью отключаем вывод Gin
	gin.DefaultWriter = io.Discard
	// gin.DefaultErrorWriter = io.Discard

	r := gin.New()

	r.POST("/users", handlers.Signup)                // ready
	r.POST("/users/{id}/notes", handlers.CreateNote) // TODO

	Log.Info("Server started", "port =", os.Getenv("PORT"))

	err := r.Run()
	if err != nil {
		// log.Error("error at starting server")
	}
}
