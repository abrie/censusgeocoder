package main

import "fmt"

import (
	"github.com/abrie/censusgeocoder"
)

func main() {
	fmt.Printf("Searching...\n")
	censusgeocoder.SearchOneLineAddress("123 Fake St")
}
