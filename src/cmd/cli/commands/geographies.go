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

func SearchGeographies(args []string) {
	flagSet := flag.NewFlagSet("geographies", flag.ExitOnError)

	benchmark := flagSet.String("benchmark", "", "Benchmark for geographies search")
	vintage := flagSet.String("vintage", "", "Vintage corresponding to Benchmark")
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

	if *vintage == "" {
		flagSet.PrintDefaults()
		fmt.Printf("\nVintage is a required parameter.\n")
		os.Exit(2)
	}

	if *onelineAddress != "" {
		searchOneLineAddressGeographies(*onelineAddress, *benchmark, *vintage)
		os.Exit(0)
	}

	if *street != "" && *city != "" && *state != "" {
		searchAddressGeographies(*street, *city, *state, *benchmark, *vintage)
		os.Exit(0)
	}

	flagSet.PrintDefaults()
	os.Exit(2)

}

func searchOneLineAddressGeographies(oneline, benchmark, vintage string) {
	result, err := censusgeocoder.SearchOneLineAddressGeographies(context.Background(), oneline, benchmark, vintage)
	if err != nil {
		log.Fatal(err)
	}

	utils.PrettyPrint(result)
}

func searchAddressGeographies(street, city, state, benchmark, vintage string) {
	result, err := censusgeocoder.SearchAddressGeographies(context.Background(), street, city, state, benchmark, vintage)
	if err != nil {
		log.Fatal(err)
	}

	utils.PrettyPrint(result)
}
