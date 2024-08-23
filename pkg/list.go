package pkg

import (
	"cligo/storage"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strings"
	"time"
)

func AddItem(item string) error {
	insertSQL := `INSERT INTO shopping_list_items (item) VALUES (?)`

	_, err := storage.DB.Exec(insertSQL, item)
	if err != nil {
		return fmt.Errorf("erreur lors de l'ajout de l'article: %w", err)
	}

	fmt.Printf("L'article '%s' a été ajouté à la base de données.\n", item)
	return nil
}

func ShowItems() {
	query := `SELECT id, item, quantity_owned, quantity_required, is_marked, last_edit, added_at FROM shopping_list_items`

	rows, err := storage.DB.Query(query)
	if err != nil {
		log.Fatalf("Erreur lors de la récupération des éléments: %v", err)
	}
	defer rows.Close()

	fmt.Printf("%-5s %-20s %-15s %-20s %-10s %-20s %-20s\n", "ID", "Item", "Owned", "Required", "Marked", "Last Edit", "Added At")
	fmt.Println(strings.Repeat("-", 100))

	for rows.Next() {
		var id int
		var item string
		var quantityOwned, quantityRequired int
		var isMarked bool
		var lastEdit, addedAt string

		err := rows.Scan(&id, &item, &quantityOwned, &quantityRequired, &isMarked, &lastEdit, &addedAt)
		if err != nil {
			log.Fatalf("Erreur lors de la lecture des résultats: %v", err)
		}

		// Convertir les chaînes de caractères en time.Time en utilisant le format ISO 8601
		lastEditTime, err := time.Parse(time.RFC3339, lastEdit)
		if err != nil {
			log.Fatalf("Erreur lors de la conversion de la date 'last_edit': %v", err)
		}

		addedAtTime, err := time.Parse(time.RFC3339, addedAt)
		if err != nil {
			log.Fatalf("Erreur lors de la conversion de la date 'added_at': %v", err)
		}

		// Formater les dates
		lastEditFormatted := lastEditTime.Format("02/01/2006")
		addedAtFormatted := addedAtTime.Format("02/01/2006")

		fmt.Printf("%-5d %-20s %-15d %-20d %-10v %-20s %-20s\n", id, item, quantityOwned, quantityRequired, isMarked, lastEditFormatted, addedAtFormatted)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Erreur lors de la lecture des résultats: %v", err)
	}
}

func MarkItem(index int) error {
	updateSQL := `UPDATE shopping_list_items 
                   SET is_marked = TRUE, last_edit = CURRENT_TIMESTAMP 
                   WHERE id = ?`

	result, err := storage.DB.Exec(updateSQL, index)
	if err != nil {
		return fmt.Errorf("erreur lors de la mise à jour de l'article: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erreur lors de la vérification des lignes affectées: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("aucun article trouvé avec l'id %d", index)
	}

	fmt.Printf("L'article avec l'id %d a été marqué.\n", index)
	return nil
}

func UnmarkItem(index int) error {
	updateSQL := `UPDATE shopping_list_items 
                   SET is_marked = FALSE, last_edit = CURRENT_TIMESTAMP 
                   WHERE id = ?`

	result, err := storage.DB.Exec(updateSQL, index)
	if err != nil {
		return fmt.Errorf("erreur lors de la mise à jour de l'article: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erreur lors de la vérification des lignes affectées: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("aucun article trouvé avec l'id %d", index)
	}

	fmt.Printf("L'article avec l'id %d a été dé-marqué.\n", index)
	return nil
}

func RemoveItem(index int) error {
	deleteSQL := `DELETE FROM shopping_list_items WHERE id = ?`

	result, err := storage.DB.Exec(deleteSQL, index)
	if err != nil {
		return fmt.Errorf("erreur lors de la suppression de l'article: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erreur lors de la vérification des lignes affectées: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("aucun article trouvé avec l'id %d", index)
	}

	fmt.Printf("L'article avec l'id %d a été supprimé.\n", index)
	return nil
}

func ReorderList(fromIndex, toIndex int) error {
	tx, err := storage.DB.Begin()
	if err != nil {
		return fmt.Errorf("erreur lors du début de la transaction: %w", err)
	}

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
		return fmt.Errorf("erreur lors de la mise à jour des index: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("erreur lors du commit de la transaction: %w", err)
	}

	fmt.Printf("Les éléments ont été réorganisés de l'index %d à l'index %d.\n", fromIndex, toIndex)
	return nil
}

func ModifyRequired(index, value int) error {
	updateSQL := `UPDATE shopping_list_items 
                   SET quantity_required = ?, last_edit = CURRENT_TIMESTAMP 
                   WHERE id = ?`

	// Execute the SQL statement
	_, err := storage.DB.Exec(updateSQL, value, index)
	if err != nil {
		return fmt.Errorf("error updating item: %w", err)
	}

	return nil
}

func ModifyOwned(index, value int) error {
	updateSQL := `UPDATE shopping_list_items 
                   SET quantity_owned = ?, last_edit = CURRENT_TIMESTAMP 
                   WHERE id = ?`

	// Execute the SQL statement
	_, err := storage.DB.Exec(updateSQL, value, index)
	if err != nil {
		return fmt.Errorf("error updating item: %w", err)
	}

	return nil
}

// ajouter fonction pour :
// corriger reorder
// trouver moyen de reasigner les id quand elment liste suppr
//correction lorsque ligne non trouver
