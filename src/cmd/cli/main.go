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
	listBenchmarksPtr := flag.Bool("listbenchmarks", false, "List all benchmarks.")
	listVintagesPtr := flag.Bool("listvintages", false, "List vintages associated with a benchmark.")
	addressPtr := flag.String("oneline", "", "Single line address")
	returnTypePtr := flag.String("returntype", "", "'location' or 'geographies'")
	benchmarkPtr := flag.String("benchmark", "", "Benchmark ID or name")
	vintagePtr := flag.String("vintage", "", "Vintage ID or name")

	flag.Parse()

	if *listBenchmarksPtr == true {
		listBenchmarks()
		os.Exit(0)
	}

	if *listVintagesPtr == true {
		if *benchmarkPtr == "" {
			fmt.Printf("censusgeocoder: Specify a benchmark to list vintages, for example:\n\n")
			fmt.Printf("\nExample: censusgeocoder -benchmark=\"Public_AR_Census2010\" -listvintages\n")
			os.Exit(1)
		} else {
			listVintages(*benchmarkPtr)
			os.Exit(1)
		}
	}

	if *addressPtr == "" {
		fmt.Printf("censusgeocoder: Specify an address to search, for example:\n\n")
		flag.PrintDefaults()
		fmt.Printf("\nExample: censusgeocoder -oneline=\"4600 Silver Hill Rd, Suitland, MD\"\n")
		os.Exit(1)
	}

	if *returnTypePtr != "locations" && *returnTypePtr != "geographies" {
		fmt.Printf("censusgeocoder: Specify returntype parameter, for example:\n\n")
		flag.PrintDefaults()
		fmt.Printf("\nExample: censusgeocoder -returntype=\"location\"\n")
		os.Exit(1)
	}

	if *benchmarkPtr == "" {
		fmt.Printf("censusgeocoder: Specify benchmark parameter, for example:\n\n")
		flag.PrintDefaults()
		fmt.Printf("\nExample: censusgeocoder -benchmark=\"Public_AR_Census2010\"\n")
		os.Exit(1)
	}

	if *returnTypePtr == "geographies" && *vintagePtr == "" {
		fmt.Printf("censusgeocoder: For 'geographies' return types, specify a vintage:\n\n")
		flag.PrintDefaults()
		fmt.Printf("\nExample: censusgeocoder -benchmark=\"Public_AR_Census2010\" -vintage=\"Census2010_Census2010\"\n")
		os.Exit(1)
	}

	//	searchAddressGeographies(*addressPtr, "Public_AR_Census2010", "Census2010_Census2010")
	listBenchmarks()
}

func searchAddressGeographies(address, benchmark, vintage string) {
	result, err := censusgeocoder.SearchOneLineAddressGeographies(context.Background(), address, benchmark, vintage)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%s\n", string(bytes))

}

func listVintages(benchmark string) {
	result, err := censusgeocoder.GetVintages(context.Background(), benchmark)

	if err != nil {
		log.Fatal(err)
	}

	bytes, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%s\n", string(bytes))
}

func listBenchmarks() {
	result, err := censusgeocoder.GetBenchmarks(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%s\n", string(bytes))
}
