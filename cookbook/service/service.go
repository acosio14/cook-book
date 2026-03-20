package service

import "github.com/acosio14/cook-book/cookbook/domain"

type Service interface {
	SaveRecipe(url string) error
	ViewList() ([]*domain.Recipe, error)
	ViewContent(id int) (*domain.Recipe, error)
	RemoveRecipe(id int) error
	AddNotes(id int, note string) error
}

func validateRecipe(*domain.Recipe) (bool, error) {
	return false, nil
}

func checkDuplicates(url string) (bool, error) {
	return false, nil
}
