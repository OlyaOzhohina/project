package cli

import (
	"flag"
	"fmt"
	"os"
)

// Specify allowed params, that could be passed in terminal.
func Params() (string, string) {
	urlCommand := flag.NewFlagSet("parse", flag.ExitOnError)

	urlLinkPtr := urlCommand.String("url", "", "Link to parse. (Required)")
	fromLinkPtr := urlCommand.String("from", "", "Parse from. Specify name. allowed: wiki-quotes")

	// If there are no passed params - exit.
	if len(os.Args) < 2 {
		fmt.Println("parse subcommand is required")
		os.Exit(1)
	}

	// If params are passed, but there is no parse param - exit.
	switch os.Args[1] {
	case "parse":
		if err := urlCommand.Parse(os.Args[2:]); err != nil {
			fmt.Print("error:  ")
			os.Exit(1)
		}
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	// If some argument has no value - exit.
	if urlCommand.Parsed() {
		if *urlLinkPtr == "" {
			urlCommand.PrintDefaults()
			os.Exit(1)
		}
		if *fromLinkPtr == "" {
			urlCommand.PrintDefaults()
			os.Exit(1)
		}
	}

	return *fromLinkPtr, *urlLinkPtr
}
