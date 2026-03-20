package repository

import "github.com/acosio14/cook-book/cookbook/domain"

type Storage interface {
	CreateNewRecipe(*domain.Recipe) error
	UpdateRecipe(*domain.Recipe) error
	ReadRecipe(id int) (*domain.Recipe, error)
	ReadAllRecipes() ([]*domain.Recipe, error)
	DeleteRecipe(id int) error
}
