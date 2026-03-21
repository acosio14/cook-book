package infrastructure

import "net/http"

type HTTPScraper struct {
	httpClient http.Client
}

// Need to create a constructor
func NewHTTPScraper() {

}

func (scraper HTTPScraper) Scrape(url string) (string, error) {

	return "string", nil
}
