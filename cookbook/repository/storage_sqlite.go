package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(path string) (*Repository, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	initializeSQLTable := `
	    CREATE TABLE IF NOT EXISTS Recipes (
		    ID INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			URL TEXT NOT NULL,
			Name TEXT NOT NULL,
			Ingredients TEXT NOT NULL,
			Instructions TEXT NOT NULL,
			Yield INTEGER,
			Notes TEXT,
		);
	`
	_, err = db.Exec(initializeSQLTable)
	if err != nil {
		return nil, err
	}

	return &Repository{db: db}, nil
}

// initialize sql table
func initTable() {

}
