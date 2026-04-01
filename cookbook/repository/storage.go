package repository

import "github.com/acosio14/cook-book/cookbook/domain"

type Storage interface {
	Add(recipe *domain.Recipe) error
	ReadContent(recipeID int) (*domain.Recipe, error)
	List() ([]domain.Recipe, error)
	Delete(recipeID int) error
}
