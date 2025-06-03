package catalog

import(
	"github.com/zakharova-e/iherb-scraper-server/internal/catalog/models"
	"context"
)

type ICatalogRepository interface {
	GetProductData(ctx context.Context, productId uint32) (*models.ProductData,error)
}
