package service

import (
	"github.com/OlyaOzhohina/project/parser/cmd/parser/scrapper"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/export"
)

// Specify Parse function by given from paramether.
func parse(from string, g *geziyor.Geziyor) *geziyor.Geziyor {
	switch from {
	case "wiki-quotes":
		g.Opt.ParseFunc = scrapper.WikiQuotes
		return g
	}
	return nil
}

func Scrap(url string, from string) {

	g := geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{url},
		ParseFunc: scrapper.WikiQuotes,
		Exporters: []export.Exporter{&export.JSON{}},
	})

	parse(from, g)

	g.Start()
}
