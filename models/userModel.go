package models

import "time"

type User struct {
	ID         int `gorm:"unique, primaryKey"`
	Username   string
	Password   string
	Created_at time.Time `gorm:"autoCreateTime"`
}
