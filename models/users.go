package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string
	Email string
}

// remember to add this user to the git setup
