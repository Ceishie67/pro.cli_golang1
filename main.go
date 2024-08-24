package main

import (
	"cligo/cmd"
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

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/shopping", shoppingHandler)

	fmt.Println("Serveur démarré sur le port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur: %v", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Bienvenue sur la liste de courses</h1><p>Utilisez /shopping pour voir et modifier votre liste de courses.</p>")
}

func shoppingHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		cmd.ShowItemsWeb(w)
	case "POST":
		item := r.FormValue("item")
		if err := cmd.AddItem(item); err != nil {
			http.Error(w, "Erreur lors de l'ajout de l'article", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/shopping", http.StatusSeeOther)
	default:
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}
