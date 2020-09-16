package commands

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

import (
	"github.com/abrie/censusgeocoder"
	"github.com/abrie/censusgeocoder/cmd/cli/utils"
)

func layerArray(str *string) []string {
	if str == nil {
		return []string{}
	} else {
		return strings.Split(*str, ",")
	}
}

func SearchGeographies(args []string) {
	flagSet := flag.NewFlagSet("geographies", flag.ExitOnError)

	benchmark := flagSet.String("benchmark", "", "Benchmark for geographies search")
	vintage := flagSet.String("vintage", "", "Vintage corresponding to Benchmark")
	onelineAddress := flagSet.String("oneline", "", "Address as one line; including street, city, and state.")
	street := flagSet.String("street", "", "Structure and Street component of address")
	city := flagSet.String("city", "", "City component of address")
	state := flagSet.String("state", "", "State component of address")
	layers := flagSet.String("layers", "", "Optional extra TigerWEB WMS Layers")
	coordinates := flagSet.String("coordinates", "", "Lat/Long coordinates seperated by a comma.")

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
		searchOneLineAddressGeographies(onelineAddress, benchmark, vintage, layerArray(layers))
		os.Exit(0)
	}

	if *coordinates != "" {
		x, y := utils.ParseCoordinates(coordinates)
		searchCoordinateGeographies(x, y, benchmark, vintage, layerArray(layers))
		os.Exit(0)
	}

	if *street != "" && *city != "" && *state != "" {
		searchAddressGeographies(street, city, state, benchmark, vintage, layerArray(layers))
		os.Exit(0)
	}

	flagSet.PrintDefaults()
	os.Exit(2)

}

func HelpGeographies() {
	fmt.Printf("usage: geographies [-oneline] [-street -city -state] [-benchmark] [-vintage]\n")
}

func searchOneLineAddressGeographies(oneline, benchmark, vintage *string, layers []string) {
	result, err := censusgeocoder.SearchOneLineAddressGeographies(context.Background(), *oneline, *benchmark, *vintage, layers)
	if err != nil {
		log.Fatal(err)
	}

	utils.PrettyPrint(result)
}

func searchAddressGeographies(street, city, state, benchmark, vintage *string, layers []string) {
	result, err := censusgeocoder.SearchAddressGeographies(context.Background(), *street, *city, *state, *benchmark, *vintage, layers)
	if err != nil {
		log.Fatal(err)
	}

	utils.PrettyPrint(result)
}

func searchCoordinateGeographies(x, y float64, benchmark, vintage *string, layers []string) {
	result, err := censusgeocoder.SearchCoordinateGeographies(context.Background(), x, y, *benchmark, *vintage, layers)
	if err != nil {
		log.Fatal(err)
	}

	utils.PrettyPrint(result)
}
