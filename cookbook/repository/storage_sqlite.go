package repository

import (
	"database/sql"
	"encoding/json"
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
			yield INTEGER
			embeddings TEXT,
		);
	`
	_, err = db.Exec(initializeSQLTable)
	if err != nil {
		return nil, err
	}

	return &Repository{db: db}, nil
}

func (repo *Repository) Add(recipe *domain.Recipe) error {
	insertRecipe := `
	    INSERT INTO Recipes (url, name, ingredients, instructions, yield)
		VALUES (?, ?, ?, ?, ?)
	`
	ingredients, err := json.Marshal(recipe.Ingredients)
	if err != nil {
		return err
	}
	instructions, err := json.Marshal(recipe.Instructions)
	if err != nil {
		return err
	}
	_, err = repo.db.Exec(
		insertRecipe,
		recipe.URL,
		recipe.Name,
		ingredients,
		instructions,
		recipe.Yield,
	)
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) ReadContent(recipeID int) (*domain.Recipe, error) {
	var recipe domain.Recipe
	var dataIngredients []byte
	var dataInstructions []byte

	selectRow := `
		SELECT id, url, name, ingredients, instructions, yield
		FROM Recipes WHERE id = ?
	`
	err := repo.db.QueryRow(selectRow, recipeID).Scan(
		&recipe.ID,
		&recipe.URL,
		&recipe.Name,
		&dataIngredients,
		&dataInstructions,
		&recipe.Yield,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("Recipe %d not found", recipeID)
	}
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(dataIngredients, &recipe.Ingredients)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(dataInstructions, &recipe.Instructions)
	if err != nil {
		return nil, err
	}
	return &recipe, nil
}

func (repo *Repository) List() ([]domain.Recipe, error) {

	selectAll := `SELECT id, name FROM Recipes`

	rows, err := repo.db.Query(selectAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipes []domain.Recipe
	for rows.Next() {
		var recipe domain.Recipe
		if err := rows.Scan(&recipe.ID, &recipe.Name); err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return recipes, nil
}

func (repo *Repository) Delete(recipeID int) error {
	deleteRecipe := `
	    DELETE FROM Recipes WHERE id = ?
	`
	_, err := repo.db.Exec(deleteRecipe, recipeID)
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) UpdateEmbedding(recipeID int, embedding []float32) error {

	return nil
}
