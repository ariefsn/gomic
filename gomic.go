package gomic

import (
	"strings"

	"github.com/ariefsn/gomic/models"
	"github.com/eaciit/toolkit"
	"github.com/gocolly/colly"
)

// Gomic = struct of gomic
type Gomic struct {
	Collector *colly.Collector
	FileName  string
}

func Init(fileName string) *Gomic {
	g := new(Gomic)
	g.Collector = colly.NewCollector()
	if fileName == "" {
		fileName = "dc"
	}
	g.FileName = fileName

	return g
}

func (G *Gomic) GetDC(series string) *models.ComicModel {
	url := "http://komikgratis.com/mg/detective-conan/" + series

	c := G.Collector

	comic := models.NewComicModel()
	comic.Name = "Detective Conan"
	comic.Alias = "DC"

	files := []models.ComicFileModel{}

	c.OnHTML(".img-responsive", func(e *colly.HTMLElement) {
		src := e.Attr("src")
		alt := e.Attr("alt")
		class := e.Attr("class")

		f := models.NewComicFileModel()
		f.Image = src
		f.ImageAlt = alt

		if !strings.Contains(class, "scan-page") {
			files = append(files, *f)
		}
	})

	c.OnRequest(func(req *colly.Request) {
		toolkit.Println("Visiting ... ", req.URL)
	})

	c.OnError(func(res *colly.Response, e error) {
		toolkit.Println("Error:", e, res.Request.URL, string(res.Body))
	})

	c.OnScraped(func(res *colly.Response) {
		comic.Files = files
	})

	c.Visit(url)

	return comic
}
