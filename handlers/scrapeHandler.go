package handlers

import (
	"context"
	"time"

	"github.com/AKauffy/BeerScraper/models"
	"github.com/gofiber/fiber/v2"
)

type searchHandler struct {
	repository models.BeerRepository
}

func (h *searchHandler) GetSearchBeer(ctx *fiber.Ctx) error {
	searchString := ctx.Params("searchString")
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	beers, err := h.repository.GetSearchBeer(context, searchString)

	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusAccepted).JSON(&fiber.Map{
		"status": "success",
		"data":   beers,
	})

}

func NewBeerSearchHandler(router fiber.Router, repository models.BeerRepository) {
	handler := &searchHandler{
		repository: repository,
	}

	router.Get("/:searchString", handler.GetSearchBeer)
}
