package models

import (
	"context"
)

type Beer struct {
	ID      string
	Name    string
	Brewery string
	ABV     string
	Rating  string
}

type BeerRepository interface {
	GetSearchBeer(ctx context.Context, searchString string) ([]Beer, error)
}
