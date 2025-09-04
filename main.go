package main

import (
	"io"
	"notes/inits"
	"notes/internal/handlers"
	"notes/pkg/logger"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	inits.EnvInit()
	inits.ConnectDB()
	inits.DoTablesDB()
}

func main() {
	// init my logger "loga"
	loga := logger.New()
	loga.Warn("DEV Level")

	// Полностью отключаем вывод Gin
	gin.DefaultWriter = io.Discard
	// gin.DefaultErrorWriter = io.Discard

	r := gin.New()

	r.POST("/users", handlers.Signup)                // ready
	r.POST("/users/{id}/notes", handlers.CreateNote) // TODO

	loga.Info("Server started", "port =", os.Getenv("PORT"))

	err := r.Run()
	if err != nil {
		// log.Error("error at starting server")
	}
}
