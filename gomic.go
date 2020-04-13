package gomic

import (
	"github.com/gocolly/colly"
)

// Gomic = struct of gomic
type Gomic struct {
	Collector *colly.Collector
}

// Init = initialize new gomic
func Init() *Gomic {
	g := new(Gomic)
	g.Collector = colly.NewCollector()

	return g
}
