package main

import (
	"fmt"
	"os"
)

import (
	"github.com/abrie/censusgeocoder/cmd/cli/commands"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("This is an unoffical cli interface to the Census.gov Geocoder.\n")
		fmt.Printf("See https://geocoding.geo.census.gov/ for the service it interacts with.\n\n")
		fmt.Printf("Usage:\n\n\tcensusgeocoder <command> [parameters]\n\n")
		fmt.Printf("The commands are:\n\n\tbenchmarks\n\tvintages\n\tlocation\n\tgeographies\n\n")
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
		commands.SearchLocations(args)
	case "geographies":
		commands.SearchGeographies(args)
	case "help":
		commands.GetHelp(args)
	default:
		fmt.Printf("Unknown command.\n")
		os.Exit(2)
	}

	os.Exit(0)
}
