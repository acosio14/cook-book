package domain

type Recipe struct {
	ID           int
	URL          string
	Name         string
	Ingredients  []string
	Instructions []string
	Yield        int
	Embedding    []float64
}
