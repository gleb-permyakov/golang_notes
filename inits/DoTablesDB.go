package inits

import "notes/models"

func DoTablesDB() {
	DB.AutoMigrate(&models.User{}, &models.Note{})
}
