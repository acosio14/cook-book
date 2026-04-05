package domain

type RecipeEvaluation struct {
	RecipeID   int
	Score      int // 1 - 10
	Feedback   string
	IsComplete bool // is it missing instructions, ingredient, text
}
