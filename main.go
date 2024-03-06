package main

import (
	"log"
	"os"

	"github.com/dating-app-test/src/infra/auth/jwt"
	"github.com/dating-app-test/src/infra/helpers"
	"github.com/dating-app-test/src/infra/persistence/postgresql"
	"github.com/dating-app-test/src/interface/rest"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotte dn")
	}
}

func main() {
	gin.SetMode(os.Getenv("GIN_MODE"))

	// postgres connection set
	dbdriver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USERNAME")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	services, err := postgresql.New(dbdriver, user, password, port, host, dbname)
	if err != nil {
		panic(err)
	}

	// set token
	tk := jwt.NewToken()

	// set helpers
	helpers := helpers.NewHelpers()

	// set routes
	r := gin.Default()

	// set routes all interface
	rest.NewRoutes(r, tk, services, helpers)

	app_port := os.Getenv("PORT")
	if app_port == "" {
		app_port = os.Getenv("API_PORT") //localhost
	}
	log.Fatal(r.Run(":" + app_port))
}
