package usecase

import (
	"context"
	"github.com/zakharova-e/iherb-scraper-server/internal/catalog"
	"github.com/zakharova-e/iherb-scraper-server/internal/catalog/models"
)

type CatalogUsecase struct {
	catalogRepo catalog.ICatalogRepository
}

func NewCatalogUsecase(catalogRepo catalog.ICatalogRepository) *CatalogUsecase {
	return &CatalogUsecase{
		catalogRepo: catalogRepo,
	}
}

func (ctl *CatalogUsecase) GetProductData(ctx context.Context, productId uint32) (*models.ProductData, error) {
	if ctl.catalogRepo == nil {
		panic("repository not specified")
	}
	return ctl.catalogRepo.GetProductData(ctx, productId)
}
