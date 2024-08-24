package handlers

import (
	"cligo/cmd"
	"cligo/pkg"
	"net/http"
)

func ShoppingHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		cmd.ShowItemsWeb(w)
	case "POST":
		item := r.FormValue("item")
		if err := pkg.AddItem(item); err != nil {
			http.Error(w, "Erreur lors de l'ajout de l'article", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/shopping", http.StatusSeeOther)
	default:
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}
