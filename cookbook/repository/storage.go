package repository

import "github.com/acosio14/cook-book/cookbook/domain"

type Storage interface {
	Add(recipe *domain.Recipe) error
	ReadContent(recipe_id int) (*domain.Recipe, error)
	List() ([]domain.Recipe, error)
	Delete(recipe_id int) error
}
