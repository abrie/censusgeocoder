package main

import (
	"fmt"
	"os"
)

import (
	"github.com/abrie/censusgeocoder/cmd/cli/commands"
)

func getLocation(args []string) {
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Need parameters...\n")
		os.Exit(1)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "benchmarks":
		commands.GetBenchmarks(args)
	case "vintages":
		commands.GetVintages(args)
	case "location":
		getLocation(args)
	case "geographies":
		commands.SearchGeographies(args)
	default:
		fmt.Printf("Unknown command.\n")
		os.Exit(2)
	}

	os.Exit(0)

	//	searchAddressGeographies(*addressPtr, "Public_AR_Census2010", "Census2010_Census2010")
}
