package main

import (
	"log"

	"github.com/lazbord/SpotyGo/services/auth/api"
	"github.com/lazbord/SpotyGo/services/auth/database"
	"github.com/lazbord/SpotyGo/services/auth/service"
)

const connectionURI = "mongodb://mongodb:27017/"

func main() {
	db, err := database.NewAdapter(connectionURI)
	if err != nil {
		log.Fatal(err)
	}

	authService := service.NewAuthService(db)
	apiAdapter := api.NewApiAdapter(authService)

	apiAdapter.NewAPI()
}
