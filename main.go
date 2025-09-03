package main

import (
	"io"
	"log"
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

	loga := logger.New()
	loga.Info("message")

	// Полностью отключаем вывод Gin
	gin.DefaultWriter = io.Discard
	// gin.DefaultErrorWriter = io.Discard

	r := gin.New()

	r.POST("/users", handlers.Signup)                // ready
	r.POST("/users/{id}/notes", handlers.CreateNote) // TODO

	logerr := log.New(os.Stderr, "[ warn  ]", log.Ltime|log.Lshortfile)
	logerr.Print("any shit")

	// log.Info("Server started", "port", os.Getenv("PORT"))

	err := r.Run()
	if err != nil {
		// log.Error("error at starting server")
	}

}
