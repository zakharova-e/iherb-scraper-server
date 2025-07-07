package catalog

import (
	"context"
	"github.com/zakharova-e/iherb-scraper-server/internal/catalog/models"
)

type ICatalogUsecase interface {
	GetProductData(ctx context.Context, productId uint32) (*models.ProductData, error)
}
