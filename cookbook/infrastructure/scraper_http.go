package infrastructure

import (
	"log"
	"net/http"
	"strings"

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

func (scraper *HTTPScraper) Scrape(url string) ([]string, error) {

	//1. Fetch URL, get raw HTML
	resp, err := scraper.httpClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	//2. Extract raw content
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// JSON-LD
	var rawContent []string
	doc.Find(`script[type="application/ld+json"]`).Each(
		func(i int, s *goquery.Selection) {
			raw := strings.TrimSpace(s.Text())
			rawContent = append(rawContent, raw)
		},
	)

	// Fallback: CSS selectors

	//TO DO: proper return values
	return rawContent, nil
}
