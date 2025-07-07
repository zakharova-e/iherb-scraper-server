package catalog

import (
	"context"
	"github.com/zakharova-e/iherb-scraper-server/internal/catalog/models"
)

type ICatalogRepository interface {
	GetProductData(ctx context.Context, productId uint32) (*models.ProductData, error)
}
