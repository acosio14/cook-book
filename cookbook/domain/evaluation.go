package domain

type RecipeEvaluation struct {
	RecipeID   int
	Score      int
	Feedback   int
	IsComplete bool
}
