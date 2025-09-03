package main

import (
	"io"
	"log/slog"
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

	log := logger.New()

	// Полностью отключаем вывод Gin
	gin.DefaultWriter = io.Discard
	// gin.DefaultErrorWriter = io.Discard

	// Создаем Gin без дефолтных middleware
	r := gin.New()

	r.POST("/users", handlers.Signup)
	r.POST("/users/{id}/notes", handlers.CreateNote)

	log.Info("Server started", "port", os.Getenv("PORT"))

	err := r.Run()
	if err != nil {
		slog.Error("error at starting server")
	}

}
