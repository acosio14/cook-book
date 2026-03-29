package repository

import (
	"database/sql"

	"github.com/acosio14/cook-book/cookbook/domain"
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
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
			url TEXT NOT NULL,
			name TEXT NOT NULL,
			ingredients TEXT NOT NULL,
			instructions TEXT NOT NULL,
			yield INTEGER,
			notes TEXT,
		);
	`
	_, err = db.Exec(initializeSQLTable)
	if err != nil {
		return nil, err
	}

	return &Repository{db: db}, nil
}

func (repo *Repository) Add(recipe *domain.Recipe) error {

	return nil
}

func (repo *Repository) Update(recipe *domain.Recipe) error {

	return nil
}

func (repo *Repository) ReadContent(recipe_id []int) (*domain.Recipe, error) {

	return nil, nil
}

func (repo *Repository) List() ([]*domain.Recipe, error) {

	return nil, nil
}

func (repo *Repository) Delete(recipe_id []int) error {

	return nil
}
