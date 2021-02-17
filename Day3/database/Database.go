package database

import (
	"SoftwareGoDay3/models"
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB      *gorm.DB
	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_PORT string
	DB_NAME string
}

func (Db *Database) LoadEnvironment() error {
	Db.DB = nil
	Db.DB_USER = os.Getenv("DB_USER")
	Db.DB_PASS = os.Getenv("DB_PASS")
	Db.DB_HOST = os.Getenv("DB_HOST")
	Db.DB_PORT = os.Getenv("DB_PORT")
	Db.DB_NAME = os.Getenv("DB_NAME")
	if Db.DB_USER == "" || Db.DB_PASS == "" || Db.DB_HOST == "" || Db.DB_PORT == "" || Db.DB_NAME == "" {
		return errors.New("Error one of the env var doesn't exist")
	}
	return nil
}

func (Db *Database) Init() error {
	if err := Db.LoadEnvironment(); err != nil {
		return err
	}
	db_str := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		Db.DB_HOST, Db.DB_PORT, Db.DB_USER, Db.DB_PASS, Db.DB_NAME)
	datab, err := gorm.Open(postgres.Open(db_str), &gorm.Config{})
	Db.DB = datab
	Db.DB.AutoMigrate(&models.Developer{})
	return err
}
