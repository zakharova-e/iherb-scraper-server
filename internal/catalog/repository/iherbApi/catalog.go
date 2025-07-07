package iherbApi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/zakharova-e/iherb-scraper-server/internal/catalog/models"
	"github.com/zakharova-e/iherb-scraper-server/internal/config"
)

type IherbApiRepository struct {
}

func NewIherbApiRepository() *IherbApiRepository {
	return &IherbApiRepository{}
}

func (repo IherbApiRepository) GetProductData(ctx context.Context, productId uint32) (data *models.ProductData, err error) {
	productDataCollector := NewIherbCollector(setDefaultProductDataCollectorHeaders)
	data = new(models.ProductData)
	productDataCollector.OnResponse(func(r *colly.Response) {
		err = json.Unmarshal(r.Body, data)
	})
	productDataCollector.OnError(func(r *colly.Response, resErr error) {
		err = resErr
	})
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occured: %w", r)
		}
	}()
	productDataCollector.Visit(config.ProductPageUrl + fmt.Sprint(productId))
	productDataCollector.Wait()
	return
}
