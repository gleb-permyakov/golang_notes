package main

import (
	"notes/handlers"
	"notes/inits"

	"github.com/gin-gonic/gin"
)

func init() {
	inits.EnvInit()
	inits.ConnectDB()
	inits.DoTablesDB()
}

func main() {
	r := gin.Default()
	r.POST("/users", handlers.Signup)
	r.Run()
}
