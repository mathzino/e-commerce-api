package main

import (
	"fp/config"
	"fp/docs"
	"fp/routes"
	"fp/utils"
	"log"

	"github.com/joho/godotenv"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {
    // for load godotenv
    // for env
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // programmatically set swagger info
    docs.SwaggerInfo.Title = "Swagger Example API"
    docs.SwaggerInfo.Description = "This is a sample server Movie."
    docs.SwaggerInfo.Version = "1.0"
    docs.SwaggerInfo.Host = utils.Getenv("SWAGGER_HOST", "localhost:8080")
    docs.SwaggerInfo.Schemes = []string{"http", "https"}

    // database connection
    db := config.ConnectDatabase()
    sqlDB, _ := db.DB()
    defer sqlDB.Close()

    // router
    r := routes.SetupRouter(db)
    // just remove port 8080
    r.Run()
}