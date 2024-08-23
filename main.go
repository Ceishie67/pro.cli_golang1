package main

import (
	"cligo/cmd"
	"cligo/storage"
	"log"
)

func main() {
	var err error
	storage.DB, err = storage.InitDB()
	if err != nil {
		log.Fatalf("Erreur d'initialisation de la base de données: %v", err)
	}

	cmd.DB = storage.DB
	cmd.Shopping()
}
