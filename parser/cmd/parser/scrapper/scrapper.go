package scrapper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

// Parse from Wiki quotes web page.
func WikiQuotes(g *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find("div.poem").Each(func(i int, s *goquery.Selection) {

		g.Exports <- s.Find("p").Text()
	})
}
