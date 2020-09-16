package commands

import (
	"fmt"
	"os"
)

func GetHelp(args []string) {
	if len(args) == 0 {
		HelpHelp()
	}

	switch args[0] {
	case "benchmarks":
		HelpBenchmarks()
	case "vintages":
		HelpVintages()
	case "locations":
		HelpLocations()
	case "geographies":
		HelpGeographies()
	}
}

func HelpHelp() {
	fmt.Printf("usage: help <command>\n")
	os.Exit(2)
}
