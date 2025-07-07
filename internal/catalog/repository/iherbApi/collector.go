package iherbApi

import (
	"github.com/gocolly/colly"
	"github.com/zakharova-e/iherb-scraper-server/internal/config"
)

type iherbCollector struct {
	*colly.Collector
}

type option func(*iherbCollector)

func NewIherbCollector(options ...option) *iherbCollector {
	c := iherbCollector{
		colly.NewCollector(),
	}
	for _, o := range options {
		o(&c)
	}
	return &c
}

func setDefaultProductDataCollectorHeaders(collector *iherbCollector) {
	if collector == nil {
		return
	}
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
