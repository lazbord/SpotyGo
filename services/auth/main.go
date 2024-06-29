package main

import (
	"log"

	"github.com/lazbord/SpotyGo/services/auth/database"
)

const connectionURI = "mongodb://localhost:27017"

func main() {
	db, err := database.NewAdapter(connectionURI)
	if err != nil {
		log.Fatal(err)
	}

	database.CreateUser(db)
}
