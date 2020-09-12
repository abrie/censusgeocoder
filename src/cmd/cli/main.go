package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

import (
	"github.com/abrie/censusgeocoder"
)

func main() {
	addressPtr := flag.String("address", "", "Single line address")
	flag.Parse()

	if *addressPtr == "" {
		fmt.Printf("censusgeocoder: Specify an address to search, for example:\n\n")
		flag.PrintDefaults()
		fmt.Printf("\nExample: censusgeocoder -address=\"4600 Silver Hill Rd, Suitland, MD\"\n")
		os.Exit(1)
	}

	fmt.Printf("Searching...")
	result, err := censusgeocoder.SearchOneLineAddressGeographies(context.Background(), *addressPtr, "Public_AR_Census2010", "Census2010_Census2010")
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%s\n", string(bytes))
}
