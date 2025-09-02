package models

import "time"

type Note struct {
	ID        int `gorm:"unique, primaryKey"`
	User_id   int
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User `gorm:"foreignKey:User_id;references:ID"`
}
