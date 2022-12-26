package infraestructure

import (
	colly2 "github.com/gocolly/colly/v2"
)

func NewFirefoxCollector() *colly2.Collector {
	collector := colly2.NewCollector()
	collector.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:108.0) Gecko/20100101 Firefox/108.0"
	return collector
}
