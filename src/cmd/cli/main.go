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
		fmt.Printf("\nExample: censusgeocoder -address=\"100 State St Chicago IL\"\n")
		os.Exit(1)
	}

	fmt.Printf("Searching...")
	result, err := censusgeocoder.SearchOneLineAddress(context.Background(), *addressPtr)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%s\n", string(bytes))
}
