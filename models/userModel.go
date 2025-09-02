package models

type User struct {
	ID         int `gorm:"unique, primaryKey"`
	Username   string
	Created_at string
}
