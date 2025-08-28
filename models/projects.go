// Package models
package models

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name    string
	Type    string
	URL     string
	GitRepo string
}
