package main

import (
	"log"

	"github.com/lazbord/SpotyGo/services/streaming/api"
	"github.com/lazbord/SpotyGo/services/streaming/database"
	"github.com/lazbord/SpotyGo/services/streaming/service"
)

const connectionURI = "mongodb://mongodb:27017/"

func main() {
	db, err := database.NewAdapter(connectionURI)
	if err != nil {
		log.Fatal(err)
	}

	authService := service.NewStreamingService(db)
	apiAdapter := api.NewApiAdapter(authService)

	apiAdapter.NewAPI()
}
