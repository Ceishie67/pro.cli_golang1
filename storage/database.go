package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	dbName := "db.sqlite"

	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		fmt.Println("Le fichier de la base de données n'existe pas encore. Création en cours...")
	} else if err != nil {
		return nil, err
	} else {
		fmt.Println("Le fichier de la base de données existe déjà.")
	}

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS shopping_list_items (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        item TEXT NOT NULL,
        quantity_owned INTEGER DEFAULT 0,
        quantity_required INTEGER NOT NULL DEFAULT 1,
        is_marked BOOLEAN DEFAULT FALSE,
        last_edit DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        added_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
    );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, fmt.Errorf("Erreur lors de la création de la table: %w", err)
	}

	fmt.Printf("La table a été vérifiée et créée avec succès dans le fichier de base de données '%s'.\n", dbName)

	return db, nil
}
