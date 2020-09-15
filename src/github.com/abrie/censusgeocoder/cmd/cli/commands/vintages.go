package commands

import (
	"context"
	"flag"
	"log"
	"os"
)

import (
	"github.com/abrie/censusgeocoder"
	"github.com/abrie/censusgeocoder/cmd/cli/utils"
)

func GetVintages(args []string) {
	flagSet := flag.NewFlagSet("vintages", flag.ExitOnError)
	benchmark := flagSet.String("benchmark", "", "List vintages of this benchmark.")

	if err := flagSet.Parse(args); err != nil {
		flagSet.PrintDefaults()
		os.Exit(2)
	}

	if *benchmark == "" {
		flagSet.PrintDefaults()
		os.Exit(2)
	}

	result, err := censusgeocoder.GetVintages(context.Background(), *benchmark)
	if err != nil {
		log.Fatal(err)
	}

	utils.PrettyPrint(result)

	os.Exit(0)
}
