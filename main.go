package main

import (
	"cligo/cmd"
	"cligo/handlers"
	"cligo/storage"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var err error
	storage.DB, err = storage.InitDB()
	if err != nil {
		log.Fatalf("Erreur d'initialisation de la base de données: %v", err)
	}

	cmd.DB = storage.DB

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/shopping", handlers.ShoppingHandler)
	http.HandleFunc("/mark_unmark", handlers.MarkUnmarkHandler)
	http.HandleFunc("/unmark_item", cmd.UnmarkItemWeb)
	http.HandleFunc("/remove_item", cmd.RemoveItemWeb)
	http.HandleFunc("/reorder_list", cmd.ReorderListWeb)
	http.HandleFunc("/modify_required", cmd.ModifyRequiredWeb)
	http.HandleFunc("/modify_owned", cmd.ModifyOwnedWeb)
	//http.HandleFunc("/reassign_ids", cmd.ReassignIDsWeb)

	fmt.Println("Serveur démarré sur le port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur: %v", err)
	}
}
