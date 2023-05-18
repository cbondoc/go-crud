package models

import "gorm.io/gorm"

// migrator database, it creates the table
type Post struct {
	gorm.Model
	Title string
	Body  string
}
