package cmd

import (
	"fmt"
	"net/http"
	"strconv"
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
func MarkItemWeb(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'ID de l'élément depuis les paramètres de l'URL
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "L'ID de l'article est manquant", http.StatusBadRequest)
		return
	}

	// Convertir l'ID en entier
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	// Vérifier l'action (marquer ou dé-marquer)
	action := r.URL.Query().Get("action")
	if action == "" {
		http.Error(w, "L'action est manquante", http.StatusBadRequest)
		return
	}

	var updateSQL string
	if action == "mark" {
		updateSQL = `UPDATE shopping_list_items SET is_marked = TRUE, last_edit = CURRENT_TIMESTAMP WHERE id = ?`
	} else if action == "unmark" {
		updateSQL = `UPDATE shopping_list_items SET is_marked = FALSE, last_edit = CURRENT_TIMESTAMP WHERE id = ?`
	} else {
		http.Error(w, "Action non reconnue", http.StatusBadRequest)
		return
	}

	// Exécuter la mise à jour dans la base de données
	result, err := DB.Exec(updateSQL, id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la mise à jour de l'article: %v", err), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la vérification des lignes affectées: %v", err), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, fmt.Sprintf("Aucun article trouvé avec l'ID %d", id), http.StatusNotFound)
		return
	}

	// Confirmation du succès
	fmt.Fprintf(w, "L'article avec l'ID %d a été %s avec succès.\n", id, action)
}

func UnmarkItemWeb(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'ID de l'élément depuis les paramètres de l'URL
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "L'ID de l'article est manquant", http.StatusBadRequest)
		return
	}

	// Convertir l'ID en entier
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	// Requête SQL pour dé-marquer l'élément
	updateSQL := `UPDATE shopping_list_items SET is_marked = FALSE, last_edit = CURRENT_TIMESTAMP WHERE id = ?`

	// Exécuter la mise à jour dans la base de données
	result, err := DB.Exec(updateSQL, id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la mise à jour de l'article: %v", err), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la vérification des lignes affectées: %v", err), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, fmt.Sprintf("Aucun article trouvé avec l'ID %d", id), http.StatusNotFound)
		return
	}

	// Confirmation du succès
	fmt.Fprintf(w, "L'article avec l'ID %d a été dé-marqué avec succès.\n", id)
}

// RemoveItemWeb handles removing an item via a web interface.
func RemoveItemWeb(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'ID de l'élément depuis les paramètres de l'URL
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "L'ID de l'article est manquant", http.StatusBadRequest)
		return
	}

	// Convertir l'ID en entier
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	// Requête SQL pour supprimer l'élément
	deleteSQL := `DELETE FROM shopping_list_items WHERE id = ?`

	// Exécuter la suppression dans la base de données
	result, err := DB.Exec(deleteSQL, id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la suppression de l'article: %v", err), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la vérification des lignes affectées: %v", err), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, fmt.Sprintf("Aucun article trouvé avec l'ID %d", id), http.StatusNotFound)
		return
	}

	// Confirmation du succès
	fmt.Fprintf(w, "L'article avec l'ID %d a été supprimé avec succès.\n", id)
}

func ReorderListWeb(w http.ResponseWriter, r *http.Request) {
	// Récupérer les indices de départ et de destination depuis les paramètres de l'URL
	fromIndexStr := r.URL.Query().Get("from")
	toIndexStr := r.URL.Query().Get("to")
	if fromIndexStr == "" || toIndexStr == "" {
		http.Error(w, "Les indices de départ et de destination sont manquants", http.StatusBadRequest)
		return
	}

	// Convertir les indices en entiers
	fromIndex, err := strconv.Atoi(fromIndexStr)
	if err != nil {
		http.Error(w, "Indice de départ invalide", http.StatusBadRequest)
		return
	}
	toIndex, err := strconv.Atoi(toIndexStr)
	if err != nil {
		http.Error(w, "Indice de destination invalide", http.StatusBadRequest)
		return
	}

	// Démarrer une transaction pour la réorganisation des éléments
	tx, err := DB.Begin()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors du début de la transaction: %v", err), http.StatusInternalServerError)
		return
	}

	// Requête SQL pour mettre à jour les indices des éléments
	updateOrderSQL := `
        UPDATE shopping_list_items
        SET id = CASE
            WHEN id = ? THEN ?
            WHEN id = ? THEN ?
            ELSE id
        END
        WHERE id IN (?, ?)`

	_, err = tx.Exec(updateOrderSQL, fromIndex, toIndex, toIndex, fromIndex, fromIndex, toIndex)
	if err != nil {
		tx.Rollback()
		http.Error(w, fmt.Sprintf("Erreur lors de la mise à jour des indices: %v", err), http.StatusInternalServerError)
		return
	}

	// Commit de la transaction
	if err := tx.Commit(); err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors du commit de la transaction: %v", err), http.StatusInternalServerError)
		return
	}

	// Confirmation du succès
	fmt.Fprintf(w, "Les éléments ont été réorganisés de l'indice %d à l'indice %d avec succès.\n", fromIndex, toIndex)
}

func ModifyRequiredWeb(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'ID de l'article et la nouvelle quantité requise depuis les paramètres de l'URL
	indexStr := r.URL.Query().Get("id")
	valueStr := r.URL.Query().Get("value")
	if indexStr == "" || valueStr == "" {
		http.Error(w, "L'ID de l'article ou la nouvelle quantité requise est manquante", http.StatusBadRequest)
		return
	}

	// Convertir l'ID et la quantité en entiers
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		http.Error(w, "ID d'article invalide", http.StatusBadRequest)
		return
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		http.Error(w, "Quantité requise invalide", http.StatusBadRequest)
		return
	}

	// Requête SQL pour mettre à jour la quantité requise de l'article
	updateSQL := `UPDATE shopping_list_items SET quantity_required = ? WHERE id = ?`

	// Exécuter la mise à jour dans la base de données
	result, err := DB.Exec(updateSQL, value, index)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la mise à jour de la quantité requise: %v", err), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la vérification des lignes affectées: %v", err), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, fmt.Sprintf("Aucun article trouvé avec l'ID %d", index), http.StatusNotFound)
		return
	}

	// Confirmation du succès
	fmt.Fprintf(w, "La quantité requise pour l'article avec l'ID %d a été modifiée à %d avec succès.\n", index, value)
}

func ModifyOwnedWeb(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'ID de l'article et la nouvelle quantité possédée depuis les paramètres de l'URL
	indexStr := r.URL.Query().Get("id")
	valueStr := r.URL.Query().Get("value")
	if indexStr == "" || valueStr == "" {
		http.Error(w, "L'ID de l'article ou la nouvelle quantité possédée est manquante", http.StatusBadRequest)
		return
	}

	// Convertir l'ID et la quantité en entiers
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		http.Error(w, "ID d'article invalide", http.StatusBadRequest)
		return
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		http.Error(w, "Quantité possédée invalide", http.StatusBadRequest)
		return
	}

	// Requête SQL pour mettre à jour la quantité possédée de l'article
	updateSQL := `UPDATE shopping_list_items SET quantity_owned = ? WHERE id = ?`

	// Exécuter la mise à jour dans la base de données
	result, err := DB.Exec(updateSQL, value, index)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la mise à jour de la quantité possédée: %v", err), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la vérification des lignes affectées: %v", err), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, fmt.Sprintf("Aucun article trouvé avec l'ID %d", index), http.StatusNotFound)
		return
	}

	// Confirmation du succès
	fmt.Fprintf(w, "La quantité possédée pour l'article avec l'ID %d a été modifiée à %d avec succès.\n", index, value)
}

//func ReassignIDsWeb() error {
//	// Démarrer une transaction pour la réattribution des IDs
//	tx, err := DB.Begin()
//	if err != nil {
//		return fmt.Errorf("erreur lors du début de la transaction: %w", err)
//	}
//
//	// Obtenir tous les articles avec leurs IDs
//	rows, err := tx.Query(`SELECT id FROM shopping_list_items ORDER BY id`)
//	if err != nil {
//		tx.Rollback()
//		return fmt.Errorf("erreur lors de la récupération des articles: %w", err)
//	}
//	defer rows.Close()
//
//	var _ int
//	newID := 1
//
//	// Réattribuer les IDs
//	for rows.Next() {
//		var id int
//		err := rows.Scan(&id)
//		if err != nil {
//			tx.Rollback()
//			return fmt.Errorf("erreur lors de la lecture des résultats: %w", err)
//		}
//
//		if id != newID {
//			_, err := tx.Exec(`UPDATE shopping_list_items SET id = ? WHERE id = ?`, newID, id)
//			if err != nil {
//				tx.Rollback()
//				return fmt.Errorf("erreur lors de la mise à jour des IDs: %w", err)
//			}
//		}
//		newID++
//	}
//
//	// Vérifier les erreurs de la boucle de résultats
//	if err := rows.Err(); err != nil {
//		tx.Rollback()
//		return fmt.Errorf("erreur lors de la lecture des résultats: %w", err)
//	}
//
//	// Commit de la transaction
//	if err := tx.Commit(); err != nil {
//		return fmt.Errorf("erreur lors du commit de la transaction: %w", err)
//	}
//
//	return nil
//}
