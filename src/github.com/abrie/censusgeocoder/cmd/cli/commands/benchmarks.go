package commands

import (
	"context"
	"fmt"
	"log"
	"os"
)

import (
	"github.com/abrie/censusgeocoder"
	"github.com/abrie/censusgeocoder/cmd/cli/utils"
)

func GetBenchmarks(args []string) {
	result, err := censusgeocoder.GetBenchmarks(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	utils.PrettyPrint(result)

	os.Exit(0)
}

func HelpBenchmarks() {
	fmt.Printf("usage: censusgeocoder benchmarks\n\n")
	fmt.Printf("Lists all available benchmarks. A benchmark is a numerical ID or\nname that references what version of the locator should be searched. This\ngenerally corresponds to MTDB data which is benchmarked twice yearly. A full list\nof options can be accessed at https://geocoding.geo.census.gov/geocoder/benchmarks.\nThe general format of the name is DatasetType_SpatialBenchmark.")
}
