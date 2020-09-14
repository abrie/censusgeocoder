package commands

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
)

import (
	"github.com/abrie/censusgeocoder"
	"github.com/abrie/censusgeocoder/cmd/cli/utils"
)

func SearchLocations(args []string) {
	flagSet := flag.NewFlagSet("location", flag.ExitOnError)

	benchmark := flagSet.String("benchmark", "", "Benchmark for geographies search")
	onelineAddress := flagSet.String("oneline", "", "Address as one line; including street, city, and state.")
	street := flagSet.String("street", "", "Structure and Street component of address")
	city := flagSet.String("city", "", "City component of address")
	state := flagSet.String("state", "", "State component of address")

	if err := flagSet.Parse(args); err != nil {
		flagSet.PrintDefaults()
		os.Exit(2)
	}

	if *benchmark == "" {
		flagSet.PrintDefaults()
		fmt.Printf("\nBenchmark is a required parameter.\n")
		os.Exit(2)
	}

	if *onelineAddress != "" {
		searchOneLineAddressLocations(*onelineAddress, *benchmark)
		os.Exit(2)
	}

	if *street != "" && *city != "" && *state != "" {
		searchAddressLocations(*street, *city, *state, *benchmark)
		os.Exit(2)
	}

	flagSet.PrintDefaults()
	os.Exit(2)

}

func searchOneLineAddressLocations(oneline, benchmark string) {
	result, err := censusgeocoder.SearchOneLineAddressLocations(context.Background(), oneline, benchmark)
	if err != nil {
		log.Fatal(err)
	}

	utils.PrettyPrint(result)
}

func searchAddressLocations(street, city, state, benchmark string) {
	result, err := censusgeocoder.SearchAddressLocations(context.Background(), street, city, state, benchmark)
	if err != nil {
		log.Fatal(err)
	}

	utils.PrettyPrint(result)
}
