package main

import (
	"fmt"
	"log"
)

import (
	"github.com/abrie/censusgeocoder"
)

func main() {
	fmt.Printf("Searching...\n")
	result, err := censusgeocoder.SearchOneLineAddress("123 Fake St.")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", result)
}
