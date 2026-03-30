package repository

import (
	"database/sql"
	"fmt"

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
	insert_recipe := `
	    INSERT INTO Recipes (url, name, ingredients, instructions, yield, notes)
		VALUES (:url, :name, :ingredients, :instructions, :yield, :notes)
	`
	_, err := repo.db.Exec(
		insert_recipe,
		recipe.ID,
		recipe.URL,
		recipe.Name,
		recipe.Ingredients,
		recipe.Instructions,
		recipe.Yield,
	)
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) ReadContent(recipe_id int) (domain.Recipe, error) {
	var recipe domain.Recipe
	select_row := `
		SELECT * FROM Recipes WHERE id = ?
	`
	err := repo.db.QueryRow(select_row, recipe_id).Scan(
		&recipe.ID,
		&recipe.URL,
		&recipe.Name,
		&recipe.Ingredients,
		&recipe.Instructions,
		&recipe.Yield,
	)
	if err == sql.ErrNoRows {
		return domain.Recipe{}, fmt.Errorf("Recipe %d not found", recipe_id)
	}
	if err != nil {
		return domain.Recipe{}, err
	}

	return recipe, nil
}

func (repo *Repository) List() ([]domain.Recipe, error) {

	select_all := `SELECT name FROM Recipes`

	rows, err := repo.db.Query(select_all)
	if err != nil {
		return nil, err
	}

	var recipes []domain.Recipe
	for rows.Next() {
		var recipe domain.Recipe
		if err := rows.Scan(&recipe.Name); err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return recipes, nil
}

func (repo *Repository) Delete(recipe_id int) error {
	deleteRecipe := `
	    DELETE FROM Recipes WHERE id = ?
	`
	_, err := repo.db.Exec(deleteRecipe, recipe_id)
	if err != nil {
		return err
	}

	return nil
}
