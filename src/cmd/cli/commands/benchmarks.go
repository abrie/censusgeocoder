package commands

import (
	"context"
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
