package infrastructure

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type HTTPScraper struct {
	httpClient *http.Client
}

// Need to create a constructor
func NewHTTPScraper() *HTTPScraper {
	return &HTTPScraper{
		httpClient: &http.Client{Timeout: 30},
	}
}

func (scraper *HTTPScraper) Scrape(url string) (string, error) { //should string resp be a http.Respones? What type would that be?
	//TO DO: Implementation Body
	//1. Fetch URL, get raw HTML
	resp, err := scraper.httpClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	//2. Extract raw recipe content
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// JSON-LD, schema.org recipe standard
	// TO-DO: Want to extract recipe information:
	//        name, recipeInstructions, recipeIngredients
	//        Package them in a structured format that
	//        service can use to generate Recipe struct.
	doc.Find()

	// Fallback: CSS selectors

	//TO DO: proper return values
	return "string", nil
}
