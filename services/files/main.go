package main

import (
	"log"

	"github.com/lazbord/SpotyGo/services/files/api"
	"github.com/lazbord/SpotyGo/services/files/database"
	"github.com/lazbord/SpotyGo/services/files/service"
)

const connectionURI = "mongodb://localhost:27017"

func main() {
	db, err := database.NewAdapter(connectionURI)
	if err != nil {
		log.Fatal(err)
	}

	FilesService := service.NewFilesService(db)
	apiAdapter := api.NewApiAdapter(FilesService)

	apiAdapter.NewAPI()
}
