package models

import "gorm.io/gorm"

type Developer struct {
	gorm.Model
	Name       string
	Age        string
	School     string
	Experience string
}
