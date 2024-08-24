package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func MarkUnmarkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Récupérer l'ID de l'article depuis les paramètres du formulaire
	idStr := r.FormValue("id")
	if idStr == "" {
		http.Error(w, "L'ID de l'article est manquant", http.StatusBadRequest)
		return
	}

	// Convertir l'ID en entier
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "ID invalide, l'ID doit être un nombre entier positif", http.StatusBadRequest)
		return
	}

	// Récupérer l'action (mark/unmark)
	action := r.FormValue("action")
	if action == "" {
		http.Error(w, "L'action est manquante", http.StatusBadRequest)
		return
	}

	// Effectuer l'action appropriée en fonction de la valeur de "action"
	switch action {
	case "mark":
		if err := cmd.M(id); err != nil {
			http.Error(w, fmt.Sprintf("Erreur lors du marquage de l'article: %v", err), http.StatusInternalServerError)
			return
		}
	case "unmark":
		if err := UnmarkItem(id); err != nil {
			http.Error(w, fmt.Sprintf("Erreur lors du dé-marquage de l'article: %v", err), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Action non reconnue, utilisez 'mark' ou 'unmark'", http.StatusBadRequest)
		return
	}

	// Redirection après succès
	http.Redirect(w, r, "/shopping", http.StatusSeeOther)
}
