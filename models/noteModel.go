package models

type Note struct {
	ID         int `gorm:"unique, primaryKey"`
	User_id    int
	Title      string
	Content    string
	Created_at string
	Updated_at string
	User       User `gorm:"foreignKey:User_id;references:ID"`
}
