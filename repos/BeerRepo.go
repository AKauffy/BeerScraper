package repos

import (
	"context"
	"fmt"
	"regexp"

	"github.com/AKauffy/BeerScraper/models"
	"github.com/gocolly/colly"
)

type BeerRepository struct{}

// GetSearchBeer implements models.BeerRepository.
func (b *BeerRepository) GetSearchBeer(ctx context.Context, searchString string) ([]models.Beer, error) {
	return scrapeBeerResults(searchString), nil
}

func scrapeBeerResults(searchString string) []models.Beer {
	scapeURL := "https://untappd.com/search?q=" + searchString

	c := colly.NewCollector(colly.AllowedDomains("https://untappd.com/", "untappd.com"))

	var beers []models.Beer

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting %s\n", r.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Printf("Error while scraping: %s\n", e.Error())
	})

	c.OnHTML("div.beer-item", func(list *colly.HTMLElement) {
		name := list.ChildText("p.name")

		brewery := list.ChildText("p.brewery")

		abv := list.ChildText("p.abv")

		rating := list.ChildText("span.num")

		beers = append(beers, models.Beer{Name: name, Brewery: brewery, ABV: abv, Rating: cleanRating(rating)})
	})

	c.OnScraped(func(r *colly.Response) {
		for _, beer := range beers {
			fmt.Printf("Name: %s, Brewery: %s, ABV: %s, rating: %s\n", beer.Name, beer.Brewery, beer.ABV, beer.Rating)
		}
	})

	c.Visit(scapeURL)
	return beers
}

func cleanRating(rating string) string {
	noParentheses := regexp.MustCompile(`[()]`).ReplaceAllString(rating, "")
	return noParentheses
}

func NewBeerRepo() models.BeerRepository {
	return &BeerRepository{}
}
