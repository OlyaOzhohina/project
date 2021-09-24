// Main implements parser.
package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
)

const parseUrl = "https://uk.wikiquote.org/wiki/%D0%A8%D0%B5%D0%B2%D1%87%D0%B5%D0%BD%D0%BA%D0%BE_%D0%A2%D0%B0%D1%80%D0%B0%D1%81_%D0%93%D1%80%D0%B8%D0%B3%D0%BE%D1%80%D0%BE%D0%B2%D0%B8%D1%87"

func parserQuotes(g *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find("div.poem").Each(func(i int, s *goquery.Selection) {

		g.Exports <- s.Find("p").Text()
	})

}

func main() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{parseUrl},
		ParseFunc: parserQuotes,
		Exporters: []export.Exporter{&export.JSON{}},
	}).Start()

}
