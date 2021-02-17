package controllers

import (
	"SoftwareGoDay3/models"
	"fmt"

	"gorm.io/gorm"
)

func CreateDeveloper(name string, age string, school string, experience string, db *gorm.DB) {
	new_db := models.Developer{Name: name, Age: age, School: school, Experience: experience}
	db.Create(&new_db)
}

func GetDevelopers(db *gorm.DB) (dev []models.Developer) {
	db.Model(&models.Developer{}).Find(&dev)
	fmt.Println(dev)
	return
}

func GetDeveloper(db *gorm.DB, id int) (dev []models.Developer) {
	db.Model(&models.Developer{}).Take(&dev, id)
	fmt.Println(dev)
	return
}
