package inits

import "notes/internal/models"

func DoTablesDB() {
	DB.AutoMigrate(&models.User{}, &models.Note{})
}
