package main

import (
	"io"
	"notes/inits"
	"notes/internal/handlers"
	"notes/internal/middleware"
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

	// make time UTC+3
	// loc := time.FixedZone("UTC+3", +3*60*60)
	// fmt.Println(time.Now().In(loc))

	// Полностью отключаем вывод Gin
	gin.DefaultWriter = io.Discard
	// gin.DefaultErrorWriter = io.Discard

	r := gin.New()

	r.POST("/users", handlers.Signup)                                                // ready
	r.POST("/users/signin", handlers.SignIn)                                         // ready
	r.POST("/users/:id/notes", middleware.CheckAuth, handlers.CreateNote)            // ready
	r.GET("/users/:id/notes", middleware.CheckAuth, handlers.GetNotes)               // ready
	r.GET("/users/:id/notes/:note_id", middleware.CheckAuth, handlers.GetNote)       // ready
	r.PUT("/users/:id/notes/:note_id", middleware.CheckAuth, handlers.PutNote)       // ready
	r.DELETE("/users/:id/notes/:note_id", middleware.CheckAuth, handlers.DeleteNote) // ready

	Log.Info("Server started", "port =", os.Getenv("PORT"))

	err := r.Run()
	if err != nil {
		// log.Error("error at starting server")
	}
}
