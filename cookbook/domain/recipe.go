package domain

type Recipe struct {
	ID           int
	URL          string
	Title        string
	Food         string
	Ingredients  []string
	Instructions []string
	Notes        []string
}
