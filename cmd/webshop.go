package cmd

import (
	"fmt"
	"net/http"
)

func ShowItemsWeb(w http.ResponseWriter) {
	query := `SELECT id, item, quantity_owned, quantity_required, is_marked FROM shopping_list_items`

	rows, err := DB.Query(query)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des éléments", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Construire la page HTML
	fmt.Fprintf(w, "<h1>Liste de Courses</h1><ul>")
	for rows.Next() {
		var id int
		var item string
		var quantityOwned, quantityRequired int
		var isMarked bool

		err := rows.Scan(&id, &item, &quantityOwned, &quantityRequired, &isMarked)
		if err != nil {
			http.Error(w, "Erreur lors de la lecture des résultats", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "<li>%d: %s (Possédé: %d, Requis: %d) - %v</li>", id, item, quantityOwned, quantityRequired, isMarked)
	}
	fmt.Fprintf(w, "</ul><form method='POST'><input type='text' name='item' placeholder='Nouvel article'/><input type='submit' value='Ajouter'/></form>")
}
