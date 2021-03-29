package main

import (
	"SoftwareGoDay2/routes"
	"SoftwareGoDay2/server"
	"log"

	"github.com/joho/godotenv"
)

//Test avec http
/*func main() {
	fmt.Printf("starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}*/

//Test avec gin
func main() {
	err := godotenv.Load("setup_env.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	router := server.NewServer()
	routes.ApplyRoutes(router.App)
	router.App.Run(":8080")
}

//Exemple donné par gin
/*func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}*/
