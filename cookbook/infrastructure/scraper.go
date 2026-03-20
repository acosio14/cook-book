package infrastructure

type Scraper interface {
	Scrape(url string) (string, error)
}
