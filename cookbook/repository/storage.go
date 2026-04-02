package repository

import "github.com/acosio14/cook-book/cookbook/domain"

type Storage interface {
	Add(recipe *domain.Recipe) error
	ReadContent(recipeID int) (*domain.Recipe, error)
	List() ([]domain.Recipe, error)
	Delete(recipeID int) error
	SaveEvaluation(evaluation *domain.RecipeEvaluation) error
	GetEvaluation(recipeID int) (*domain.RecipeEvaluation, error)
	UpdateEmbedding(recipeID int, embedding []float32) error
	SearchByEmbedding(embedding []float32, limit int) ([]domain.Recipe, error)
}
