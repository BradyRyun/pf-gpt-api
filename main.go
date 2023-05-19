//	@title			Personal Finance API
//	@version		0.1.0
//	@description	Generate a financial plan via ChatGPT
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	brady@ryunengineering.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/
//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

package main

import (
	"github.com/joho/godotenv"
	"log"
	"personal-finance-api/models"
	"personal-finance-api/router"
	"personal-finance-api/services"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	models.ConnectFirestore()
	services.InitSendGrid()
	services.InitChatGpt()
	r := router.SetupRouter()

	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server...")
	}
}
