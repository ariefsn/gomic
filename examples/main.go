package main

import (
	"github.com/ariefsn/gomic"
	"github.com/eaciit/toolkit"
)

func main() {
	g := gomic.Init()

	res := g.GetDC("004")

	toolkit.Println("res", toolkit.JsonStringIndent(res, "\n"))
}
