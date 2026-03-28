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

// NOTE: Getting every possible schema from scrapper. Need to handle extra content and only
// retrieve recipe.
type RecipeSchema struct {
	Name         string `json:"name"`
	Ingredients  string `json:"recipeIngredient"`
	Instructions string `json:"recipeInstruction"`
}

//var recipe RecipeSchema
//json.Unmarshal([]byte(raw), &recipe)

// Will recieve content from scraper and give it to AI with prompt to get exact recipe content.
// maps AI response to Recipe Struct
