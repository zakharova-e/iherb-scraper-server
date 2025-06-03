package iherbApi

import (
	"github.com/gocolly/colly"
	"github.com/zakharova-e/iherb-scraper-server/internal/catalog/models"
	"github.com/zakharova-e/iherb-scraper-server/internal/config"
	"fmt"
	"encoding/json"
	"context"
)

type iherbCollector struct{
	*colly.Collector
}

type option func(*iherbCollector)

type IherbApiRepository struct{
	collector *iherbCollector
}

func NewIherbApiRepository(collector *iherbCollector) *IherbApiRepository{
	if collector == nil{
		collector = NewIherbCollector(setDefaultCollectorHeaders)
	}
	return &IherbApiRepository{
		collector: collector,
	}
}

func NewIherbCollector(options ...option) *iherbCollector{
	c := iherbCollector{
			colly.NewCollector(),
		}
	for _,o := range options{
		o(&c)
	}

	return &c
}

func setDefaultCollectorHeaders(collector *iherbCollector){
	collector.OnRequest(func(r *colly.Request) {
        r.Headers.Set("accept-language", config.HttpDefaultAcceptLanguageHeader)
        r.Headers.Set("platform", config.HttpDefaultPlatformHeader)
		r.Headers.Set("regiontype", config.HttpDefaultRegionTypeHeader)
		r.Headers.Set("ih-pref", config.HttpDefaultIhPrefHeader)
		r.Headers.Set("pref", config.HttpDefaultPrefHeader)
		r.Headers.Set("user-agent", config.HttpDefaultUserAgentHeader)
		r.Headers.Set("content-type", config.HttpDefaultContentTypeHeader)
    })
}



func (repo IherbApiRepository) GetProductData(ctx context.Context, productId uint32) (*models.ProductData,error){
	var (
		data models.ProductData
		err error
	)
	repo.collector.OnResponse(func(r *colly.Response) {
		//fmt.Println(string(r.Body))
		err = json.Unmarshal(r.Body, &data)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}
        //fmt.Printf("Response: %+v\n", data)
	})

	repo.collector.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	repo.collector.Visit(config.ProductPageUrl+fmt.Sprint(productId)+"")
    repo.collector.Wait()
	return &data, err
}