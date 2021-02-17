package models

import (
	"log"

	"gorm.io/gorm"
)

type Developer struct {
	gorm.Model
	Name       string
	Age        string
	School     string
	Experience string
}

func (d *Developer) Delete(db *gorm.DB) error {
	if err := db.Model(&Developer{}).Delete(d).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}
