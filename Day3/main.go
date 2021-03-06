package main

import (
	"SoftwareGoDay3/controllers"
	"SoftwareGoDay3/database"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err_env := godotenv.Load("docker_env.env")
	if err_env != nil {
		log.Fatal(err_env)
	}
	db := database.Database{}
	err := db.Init()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database my_database is ready ")
	controllers.CreateDeveloper("Boug", "12", "epitech", "jesuistropfort", db.DB)
	controllers.GetDevelopers(db.DB)
	fmt.Printf("\n\n\n")
	controllers.GetDeveloper(db.DB, 3)
	fmt.Printf("\n\n\n")
	controllers.UpdateDeveloper(db.DB, 3, "Name", "revien")
	fmt.Printf("\n\n\n")
	controllers.DeleteDeveloper(db.DB, 13)
	fmt.Printf("\n\n\n")
	controllers.GetDevelopers(db.DB)
}
