package models

import (
	"context"
)

type Beer struct {
	ID      string
	Name    string
	Brewery string
	ABV     string
	rating  string
}

type BeerRepository interface {
	GetMany(ctx context.Context) ([]*Beer, error)
	GetOne(ctx context.Context, beerID string) (*Beer, error)
	CreateOne(ctx context.Context, Beer Beer) (*Beer, error)
}
