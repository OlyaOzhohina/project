// Main implements parser.
package main

import (
	"github.com/OlyaOzhohina/project/parser/cmd/parser/cli"
	"github.com/OlyaOzhohina/project/parser/cmd/parser/service"
)

func main() {
	// Parse cli paramethers.
	from, url := cli.Params()

	// Start scrapping.
	service.Scrap(url, from)
}
