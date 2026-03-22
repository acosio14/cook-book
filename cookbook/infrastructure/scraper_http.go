package infrastructure

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type HTTPScraper struct {
	httpClient *http.Client
}

// Need to create a constructor
func NewHTTPScraper() *HTTPScraper {
	return &HTTPScraper{
		httpClient: &http.Client{},
	}
}

func (scraper *HTTPScraper) Scrape(url string) (string, error) { //should string resp be a http.Respones? What type would that be?
	//TO DO: Implementation Body
	//1. Fetch URL, get raw HTML
	resp, err := scraper.httpClient.Get(url)
	if err != nil {
		// handle request failure
	}
	defer resp.Body.Close()
	//2. Extract desired content
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		// handle error
	}
	doc.Find()
	//TO DO: real error handling
	return "string", nil
}
